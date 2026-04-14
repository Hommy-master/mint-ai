package db

import (
	"context"
	"cozeos/internal/consts"
	"cozeos/internal/model"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/os/glog"
	"gorm.io/gorm"
)

func GetUserByOpenID(ctx context.Context, openID string) (*model.User, error) {
	// 从数据库中查询用户信息
	var user model.User

	err := NewDB().Where("wechat_id = ?", openID).First(&user).Error
	if err != nil {
		glog.Warningf(ctx, "query user failed, err: %+v, openID: %s", err, openID)
		return nil, fmt.Errorf("query user failed")
	}

	return &user, nil
}

func GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	// 从数据库中查询用户信息
	var user model.User

	err := NewDB().Where("id = ?", id).First(&user).Error
	if err != nil {
		glog.Warningf(ctx, "query user failed, err: %+v, id: %d", err, id)
		return nil, fmt.Errorf("query user failed")
	}

	return &user, nil
}

func CreateUser(ctx context.Context, u *model.User) error {
	err := NewDB().Create(u).Error
	if err != nil {
		glog.Errorf(ctx, "create user failed, err: %+v, user: %+v", err, *u)
		return fmt.Errorf("create user failed")
	}

	return nil
}

func UpdateUser(ctx context.Context, user *model.User, fields []string) error {
	glog.Debugf(ctx, "update user, user: %+v, fields: %+v", *user, fields)

	// 从数据库中查询用户信息
	query := NewDB().Model(&model.User{}).Where("id = ?", user.ID)
	if len(fields) > 0 {
		query = query.Select(fields)
	}
	err := query.Updates(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") &&
			strings.Contains(err.Error(), "users.idx_users_name") {
			glog.Warningf(ctx, "update user failed, unique index conflict, err: %+v", err)
			return fmt.Errorf("username already exists")
		}

		glog.Warningf(ctx, "update user failed, err: %+v, user: %+v", err, *user)
		return fmt.Errorf("update user failed")
	}

	return nil
}

// @brief 更新用户积分，需要做事务处理
// @param ctx 请求上下文信息
// @param userID 用户ID
// @param points 用户积分，如果为正，表示积分增加，为负，表示积分减少
// @param desc 积分消费描述，如：插件充值、新人奖励、视频插件调用等
// @param orderNO 订单号，仅充值时，订单号不为空，其它情况没有订单号
// @return 成功则返回nil，否则返回相应错误
func UpdateUserPoints(ctx context.Context, userID uint, points float64, desc string, orderNO string) error {
	// 1. 更新用户积分
	result := NewDB().Model(&model.User{}).Where("id = ?", userID).Update("points", gorm.Expr("points + ?", points))
	if result.Error != nil {
		glog.Warningf(ctx, "update user points failed, err: %+v, userID: %d, points: %f", result.Error, userID, points)
		return fmt.Errorf("update user points failed: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user with ID %d not found", userID)
	}

	// 2. 记录积分消费日志
	log := &model.PluginBalanceLog{
		UserID:      userID,
		Points:      points,
		OrderNO:     orderNO,
		Description: desc,
	}
	err := NewDB().Create(log).Error
	if err != nil {
		glog.Warningf(ctx, "create points log failed, err: %+v, log: %+v", err, *log)
		return fmt.Errorf("create points log failed")
	}

	return nil
}

// 检查用户积分是否足够，如果足够，则返回nil，否则返回相应错误，注意：只有用户积分足够，才会返回用户信息
func CheckUserPoints(ctx context.Context, price float64) (*model.User, error) {
	userID := ctx.Value("id").(uint)

	// 1. 从数据库中查询用户信息
	u, err := GetUserByID(ctx, userID)
	if err != nil {
		glog.Warningf(ctx, "get user failed, err: %+v, userID: %d", err, userID)
		return nil, fmt.Errorf("get user failed")
	}

	// 2. 检查用户积分是否足够
	if u.Points < price {
		glog.Warningf(ctx, "user points not enough, userID: %d, points: %.2f, price: %.2f", userID, u.Points, price)
		return nil, fmt.Errorf("user points not enough")
	}

	return u, nil
}

// @brief 创建用户，如果用户存在，就直接返回当前存在的用户，如果不存在，就创建一新用户，并返回新用户信息
// @param ctx 上下文信息
// @param user 用户信息
// @return 新用户信息，可能是已存在的用户信息，如果成功，则error为nil，否则包含相应错误提示
func CreateUserEx(ctx context.Context, user *model.User) (*model.User, error) {
	// 1. 开启事务确保原子性
	tx := NewDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 2. 尝试通过唯一字段查找现有用户
	existingUser := &model.User{}
	query := tx.Where("wechat_id = ?", user.WeChatID)
	if err := query.First(existingUser).Error; err != nil {
		// 用户不存在，继续创建流程
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			glog.Warningf(ctx, "query user failed, err: %+v, wechatID: %s", err, user.WeChatID)
			tx.Rollback()
			return nil, fmt.Errorf("query user failed")
		}
	} else {
		// 用户已存在，提交事务并返回
		tx.Commit()
		return existingUser, nil
	}

	// 3. 创建用户
	if err := tx.Create(user).Error; err != nil {
		glog.Warningf(ctx, "create user failed, err: %+v, user: %+v", err, *user)
		tx.Rollback()
		return nil, fmt.Errorf("create user failed")
	}

	// 4. 提交事务
	if err := tx.Commit().Error; err != nil {
		glog.Warningf(ctx, "commit transaction failed, err: %+v, user: %+v", err, *user)
		return nil, fmt.Errorf("commit transaction failed")
	}

	// 5. 再次查询用户信息
	newUser, err := GetUserByOpenID(ctx, user.WeChatID)
	if err != nil {
		glog.Warningf(ctx, "get user failed, err: %+v, wechatID: %s", err, user.WeChatID)
		return nil, fmt.Errorf("get user failed")
	}

	// 6. 新用户，记录新人积分增送
	AddBalanceLog(ctx, &model.PluginBalanceLog{
		UserID:      newUser.ID,
		Points:      user.Points,
		Description: consts.NewUserRewardDesc,
		OrderNO:     "",
	})

	return newUser, nil
}
