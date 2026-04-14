package user

import (
	"bytes"
	"context"
	"cozeos/internal/config"
	"cozeos/internal/consts"
	"cozeos/internal/errcode"
	"cozeos/internal/model"
	"cozeos/internal/pkg/cache"
	"cozeos/internal/pkg/db"
	"cozeos/internal/pkg/helper"
	"cozeos/internal/pkg/i"
	"cozeos/internal/types"
	"cozeos/util"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dchest/captcha"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/google/uuid"
	"github.com/silenceper/wechat/v2"
	"gorm.io/gorm"
)

// 获取随机验证码，返回一张图片，可以用HTML的<img>标签显示
func AuthCaptcha(ctx context.Context, phone string) error {
	r := g.RequestFromCtx(ctx)

	// 生成验证码图片
	id := captcha.NewLen(4)
	captcha.WriteImage(r.Response.Writer, id, 200, 80)

	r.Session.Set("phone", phone)
	r.Session.Set("captcha", id)
	return nil
}

// 验证验证码
func AuthCaptchaVerify(ctx context.Context, phone string, code string) error {
	r := g.RequestFromCtx(ctx)

	// 验证手机号是否正确
	val, err := r.Session.Get("phone")
	if err != nil || val.IsNil() {
		glog.Warningf(ctx, "get session of phone failed: %+v", err)
		return i.T(ctx, errcode.CodeCaptchaNotRequested)
	}

	if val.String() != phone {
		glog.Warningf(ctx, "phone not match, session-phone: %s, request: %s", val.String(), phone)
		return i.T(ctx, errcode.CodeInvalidPhoneNumber)
	}

	val, err = r.Session.Get("captcha")
	if err != nil || val.IsNil() {
		glog.Warningf(ctx, "get session of captcha failed: %+v", err)
		return i.T(ctx, errcode.CodeCaptchaNotRequested)
	}

	// 验证验证码
	if !captcha.VerifyString(val.String(), code) {
		glog.Debugf(ctx, "verify captcha failed, session-captcha: %s, code: %s", val.String(), code)
		return i.T(ctx, errcode.CodeInvalidCaptcha)
	}

	return nil
}

// 发送短信验证码
func AuthSms(ctx context.Context, phone string) error {
	return nil
}

// 验证短信验证码
func AuthSmsVerify(ctx context.Context, phone string, code string) error {
	return nil
}

// 用户名密码登录
func AuthLogin(ctx context.Context, phone string, pass string) (*types.User, error) {
	var u model.User

	db := db.NewDB()
	err := db.Model(&model.User{}).Where("phone = ?", phone).First(&u).Error
	if err != nil {
		glog.Warningf(ctx, "get user failed: %+v", err)

		// 用户不存在，创建一个新用户
		u.Phone = phone
		u.Pass = util.MD5(pass)
		u.Name = "User" + phone
		u.APIKey = uuid.New().String()
		err = db.Model(&model.User{}).Create(&u).Error
		if err != nil {
			glog.Errorf(ctx, "create user failed: %+v", err)
			return nil, i.T(ctx, errcode.CodeInternalServerError)
		}

		err = db.Model(&model.User{}).Where("phone =?", phone).First(&u).Error
		if err != nil {
			glog.Errorf(ctx, "get user failed: %+v", err)
			return nil, i.T(ctx, errcode.CodeInternalServerError)
		}

		glog.Infof(ctx, "login success, phone: %s, ip: %s", phone, g.RequestFromCtx(ctx).GetClientIp())
		return &types.User{ID: u.ID, Name: u.Name, Phone: u.Phone, Points: u.Points, Email: u.Email, WeChatID: u.WeChatID, APIKey: u.APIKey}, nil
	}

	if u.Pass != util.MD5(pass) {
		glog.Warningf(ctx, "password not match, phone: %s, pass: %s", phone, pass)
		return nil, i.T(ctx, errcode.CodeInvalidPassword)
	}

	glog.Infof(ctx, "login success, phone: %s, ip: %s", phone, g.RequestFromCtx(ctx).GetClientIp())
	return &types.User{ID: u.ID, Name: u.Name, Phone: u.Phone, Points: u.Points, Email: u.Email, WeChatID: u.WeChatID, APIKey: u.APIKey}, nil
}

func Check(ctx context.Context, phone string) (bool, error) {
	var u model.User

	db := db.NewDB()
	err := db.Model(&model.User{}).Where("phone =?", phone).First(&u).Error
	if err != nil {
		glog.Warningf(ctx, "get user failed: %+v", err)
		return false, nil
	}

	return true, nil
}

// UI获取验证码
func AuthQRCode(ctx context.Context) (string, int, error) {
	// 1. 获取访问token
	wc := wechat.NewWechat()
	oa := wc.GetOfficialAccount(config.GetWeChatConfig())
	token, err := oa.GetAccessToken()
	if err != nil {
		glog.Errorf(ctx, "get access token failed, err: %+v", err)
		return "", 0, i.T(ctx, errcode.CodeInternalServerError)
	}

	// 2. 生成二维码
	url := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + token
	requestBody := map[string]interface{}{
		"expire_seconds": consts.WeChatQRCodeExpire,
		"action_name":    "QR_SCENE",
		"action_info": map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_id": config.SceneID,
			},
		},
	}
	jsonBody, _ := json.Marshal(requestBody)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		glog.Errorf(ctx, "create qrcode failed, err: %+v", err)
		return "", 0, i.T(ctx, errcode.CodeCreateQRCodeFailed)
	}
	defer resp.Body.Close()

	var result struct {
		Ticket string `json:"ticket"`
		URL    string `json:"url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		glog.Errorf(ctx, "decode response failed, err: %+v", err)
		return "", 0, i.T(ctx, errcode.CodeCreateQRCodeFailed)
	}
	glog.Debugf(ctx, "ticket: %s, url: %s", result.Ticket, result.URL)

	// 3. 缓存ticket做为用户是否登录的凭证
	key := fmt.Sprintf(consts.TicketKeyFormat, result.Ticket)
	cache.GetInstance().Set(key, "", time.Duration(consts.WeChatQRCodeExpire)*time.Second)

	return result.Ticket, consts.WeChatQRCodeExpire, nil
}

// 校验二维码状态
func AuthQRCodeStatus(ctx context.Context, sceneID string, ticket string) (*types.User, error) {
	glog.Debugf(ctx, "sceneID: %s, ticket: %s", sceneID, ticket)

	// 1. 验证用户是否登录
	key := fmt.Sprintf(consts.TicketKeyFormat, ticket)
	val, bexist := cache.GetInstance().Get(key)
	if !bexist {
		glog.Warningf(ctx, "get cache failed, key: %s", key)
		return nil, i.T(ctx, errcode.CodeInvalidTicket)
	}

	openID, ok := val.(string)
	if !ok {
		glog.Warningf(ctx, "invalid cache value type, key: %s", key)
		return nil, i.T(ctx, errcode.CodeInvalidTicket)
	}

	if openID == "" {
		glog.Infof(ctx, "qrcode login pending, key: %s", key)
		return nil, i.T(ctx, errcode.CodeQRCodeLoginPending)
	}

	// 2. 登录成功，从数据库中查询用户
	var u model.User
	db := db.NewDB()
	err := db.Model(&model.User{}).Where("wechat_id =?", openID).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			glog.Warningf(ctx, "user not found, key: %s", key)
			return nil, i.T(ctx, errcode.CodeQRCodeLoginFailed)
		}

		glog.Warningf(ctx, "get user failed: %+v", err)
		return nil, i.T(ctx, errcode.CodeInternalServerError)
	}

	// 3. 返回用户信息
	return helper.ModelUserToTypesUser(&u), nil
}

// 需要检查session状态
func QueryUser(ctx context.Context, id uint) (*types.User, error) {
	var u model.User

	db := db.NewDB()
	err := db.Model(&model.User{}).Where("id =?", id).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			glog.Warningf(ctx, "user not found, id: %d", id)
			return nil, i.T(ctx, errcode.CodeUserNotFound)
		}

		glog.Warningf(ctx, "get user failed: %+v", err)
		return nil, i.T(ctx, errcode.CodeInternalServerError)
	}

	return helper.ModelUserToTypesUser(&u), nil
}

func UpdateUser(ctx context.Context, id uint, name *string, apiKey *string) error {
	fields := []string{}
	var u model.User

	u.ID = id

	if name != nil {
		u.Name = *name
		fields = append(fields, "name")
	}

	if apiKey != nil {
		u.APIKey = *apiKey
		fields = append(fields, "api_key")
	}

	if len(fields) == 0 {
		glog.Infof(ctx, "no need to update user info, id: %d", id)
		return nil
	}

	err := db.UpdateUser(ctx, &u, fields)
	if err != nil {
		// 唯一索引冲突
		if strings.Contains(err.Error(), "username already exists") {
			glog.Warningf(ctx, "update user failed, id: %d, err: %+v", id, err)
			return i.T(ctx, errcode.CodeDuplicateUserName)
		}

		glog.Warningf(ctx, "update user failed, id: %d, err: %+v", id, err)
		return i.T(ctx, errcode.CodeInternalServerError)
	}

	return nil
}

// @brief 查询用户积分
// @param ctx HTTP请求上下文
// @param apiKey 用户调用API使用的Key
// @return 积分数，错误信息
func Query(ctx context.Context, apiKey string) (float64, error) {
	usr, err := db.GetUserByAPIKey(ctx, apiKey)
	if err != nil {
		glog.Warningf(ctx, "db.GetUserByAPIKey failed, apiKey: %s, error: %v", apiKey, err)
		return 0, i.T(ctx, errcode.CodeInvalidAPIKey)
	}

	glog.Infof(ctx, "db.GetUserByAPIKey success, apiKey: %s, user: %+v", apiKey, usr)
	return usr.Points, nil
}

// @brief 减少用户积分
// @param ctx HTTP请求上下文
// @param apiKey 用户调用API使用的Key
// @param points 待扣除的积分数
// @param desc 扣除积分的理由，什么场景下扣除的积分
// @return 错误信息
func Deduct(ctx context.Context, apiKey string, points float64, desc string) error {
	// 1. 查询用户信息
	usr, err := db.GetUserByAPIKey(ctx, apiKey)
	if err != nil {
		glog.Warningf(ctx, "db.GetUserByAPIKey failed, apiKey: %s, error: %v", apiKey, err)
		return i.T(ctx, errcode.CodeInvalidAPIKey)
	}

	// 2. 根据用户VIP等级和有效期计算实际扣费
	actualPoints := points // 默认原价
	var discountDesc string
	currentTime := time.Now()
	switch usr.VIPLevel {
	case 1: // VIP用户：8折
		// 检查VIP是否在有效期内
		if usr.VIPExpireAt.After(currentTime) {
			actualPoints = points * 0.8
			discountDesc = "(VIP 8折)"
		} else {
			discountDesc = "(VIP已过期)"
		}
	case 2, 3: // SVIP用户：6折
		// 检查SVIP是否在有效期内
		if usr.VIPExpireAt.After(currentTime) {
			actualPoints = points * 0.6
			discountDesc = "(SVIP 6折)"
		} else {
			discountDesc = "(SVIP已过期)"
		}
	default:
		discountDesc = "(普通用户)"
	}

	// 3. 添加折扣信息到描述
	finalDesc := fmt.Sprintf("%s %s", desc, discountDesc)

	// 4. 扣除积分 & 增加用户消费记录
	err = db.UpdateUserPoints(ctx, usr.ID, -actualPoints, finalDesc, "")
	if err != nil {
		glog.Warningf(ctx, "Update user points failed, apiKey: %s, original points: %f, actual points: %f, desc: %s, err: %+v",
			apiKey, points, actualPoints, finalDesc, err)
		return i.T(ctx, errcode.CodeInternalServerError)
	}

	glog.Infof(ctx, "Update user points success, apiKey: %s, original points: %f, actual points: %f, desc: %s", apiKey, points, actualPoints, finalDesc)
	return nil
}
