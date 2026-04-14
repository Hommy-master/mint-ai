package user

import (
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
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"gorm.io/gorm"
)

var (
	client    *core.Client
	nativeSvc *native.NativeApiService
)

// 初始化微信支付客户端（在应用启动时调用）
func init() {
	ctx := context.Background()

	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(config.WeChatMchPrivateKeyPath)
	if err != nil {
		glog.Warningf(context.Background(), "load private key failed: %v", err)
		return
	}

	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(
			config.WeChatMchID,
			config.WeChatMchCertSN,
			mchPrivateKey,
			config.WeChatMchAPIv3Key),
	}

	client, err = core.NewClient(ctx, opts...)
	if err != nil {
		glog.Warningf(ctx, "create wechat pay client failed: %v", err)
		return
	}

	nativeSvc = &native.NativeApiService{Client: client}
}

// 微信支付 - 创建订单
func CreateOrder(ctx context.Context, payMethod string, productID string, productValue int) (string, string, int, int, error) {
	userID := ctx.Value("id").(uint)

	glog.Infof(ctx, "CreateOrder, userID: %d, payMethod: %s, product: %s", userID, payMethod, productID)

	// 1. 计算商品价格
	price, err := calculateProductPrice(ctx, productID, productValue)
	if err != nil {
		return "", "", 0, 0, err
	}

	// 2. 生成唯一订单号
	orderNo := fmt.Sprintf("WXPAY%s%s", time.Now().Format("20060102150405"), util.RandString(6)) // 最后6位是随机数

	// 3. 调用微信支付API
	resp, _, err := nativeSvc.Prepay(ctx, native.PrepayRequest{
		Appid:       core.String(config.WeChatAPPID),
		Mchid:       core.String(config.WeChatMchID),
		Description: core.String(productID),
		OutTradeNo:  core.String(orderNo),
		NotifyUrl:   core.String(config.WeChatPayNotifyURL),
		Amount: &native.Amount{
			Total: core.Int64(price),
		},
	})
	if err != nil {
		glog.Warning(ctx, "prepay failed, orderNO: %s, err: %v", orderNo, err)
		return "", "", 0, 0, i.T(ctx, errcode.CodeCreateOrderFailed)
	}

	if resp.CodeUrl == nil {
		glog.Warningf(ctx, "prepay failed, orderNO: %s, resp: %v", orderNo, resp)
		return "", "", 0, 0, i.T(ctx, errcode.CodeCreateOrderFailed)
	}

	// 4. 初始化订单状态
	key := fmt.Sprintf(consts.OrderNOKeyFormat, orderNo)
	orderInfo := types.OrderInfo{
		UserID:      userID,
		ProductID:   consts.ProductID(productID),
		Price:       price,
		OrderStatus: consts.OrderStatusNotPay,
	}
	cache.GetInstance().Set(key, orderInfo, consts.PayTimeout)

	glog.Infof(ctx, "prepay success, userID: %d, orderNO: %s, resp: %v",
		userID, orderNo, resp)
	return *resp.CodeUrl, orderNo, int(consts.PayTimeout.Seconds()), int(price / 100), nil
}

// 微信支付 - 支付回调
func WeChatNotify(ctx context.Context) {
	r := g.RequestFromCtx(ctx)

	// 1. 解析回调请求
	var req types.WeChatPayNotify
	if err := r.Parse(&req); err != nil {
		glog.Warningf(ctx, "WeChatNotify parse failed: %+v", err)
		r.Response.WriteStatusExit(400, "Parse request failed")
		return
	}
	glog.Infof(ctx, "ID:%s, CreateTime:%s,ResourceType:%s,EventType:%s,Summary:%s",
		req.ID, req.CreateTime, req.ResourceType, req.EventType, req.Summary)

	// 2. 解密回调数据
	decryptedData, err := decryptWeChatResource(
		[]byte(config.WeChatMchAPIv3Key),
		[]byte(req.Resource.AssociatedData),
		[]byte(req.Resource.Nonce),
		req.Resource.Ciphertext,
	)
	if err != nil {
		glog.Errorf(ctx, "decrypt resource failed: %v", err)
		r.Response.WriteStatusExit(400, "Decrypt failed")
		return
	}

	// 3. 解析解密后的数据
	var resource types.DecryptedResource
	if err = json.Unmarshal(decryptedData, &resource); err != nil {
		glog.Errorf(ctx, "unmarshal decrypted data failed: %v", err)
		r.Response.WriteStatusExit(400, "Invalid resource data")
		return
	}
	glog.Infof(ctx, "decrypted resource: %+v", resource)

	// 4. 处理订单状态
	if resource.TradeState == "SUCCESS" {
		orderNO := resource.OutTradeNo
		key := fmt.Sprintf(consts.OrderNOKeyFormat, orderNO)
		orderInfo, ok := cache.GetInstance().Get(key)
		if ok {
			orderInfo := orderInfo.(types.OrderInfo)
			// 统一调用支付成功处理函数
			handlePaySuccess(ctx, orderInfo, resource)
		} else {
			glog.Warningf(ctx, "order NO (%s) not found, transactionID: %s", orderNO, resource.TransactionID)
		}

		// 支付成功，无论如何都需要记录日志
		glog.Infof(ctx, "order payed successfully, orderNO: %s, total: %.2f, transactionID: %s, success time: %s",
			resource.OutTradeNo, float64(resource.Amount.Total)/100.0, resource.TransactionID, resource.SuccessTime)
	} else {
		glog.Warningf(ctx, "payment not success, trade state: %s, transactionID: %s, desc: %s",
			resource.TradeState, resource.TransactionID, resource.TradeStateDesc)
	}

	// 5. 返回成功响应
	r.Response.WriteXml(g.Map{
		"code":    "SUCCESS",
		"message": "OK",
	})
}

// 微信支付 - 查询支付结果
func CheckPayment(ctx context.Context, orderNO string) (string, error) {
	key := fmt.Sprintf(consts.OrderNOKeyFormat, orderNO)
	orderInfo, ok := cache.GetInstance().Get(key)
	if ok {
		tmp := orderInfo.(types.OrderInfo)
		// 如果本地状态是已支付，直接返回
		if tmp.OrderStatus == consts.OrderStatusPayed {
			glog.Infof(ctx, "userID: %d, orderNO: %s, order status: %s (from cache)",
				tmp.UserID, orderNO, string(tmp.OrderStatus))
			return string(tmp.OrderStatus), nil
		}
	}

	// 本地没有订单或状态未支付，主动查询微信支付
	tradeState, err := queryWeChatOrder(ctx, orderNO)
	if err != nil {
		glog.Warningf(ctx, "query wechat order failed, orderNO: %s, err: %v", orderNO, err)
		return "", err
	}

	// 根据微信支付状态更新本地
	switch tradeState {
	case "SUCCESS":
		// 如果本地缓存中有订单信息
		if ok {
			orderInfo := orderInfo.(types.OrderInfo)
			// 查询订单详情
			resp, _, err := nativeSvc.QueryOrderByOutTradeNo(ctx, native.QueryOrderByOutTradeNoRequest{
				OutTradeNo: core.String(orderNO),
				Mchid:      core.String(config.WeChatMchID),
			})
			if err != nil {
				glog.Warningf(ctx, "query order detail failed, orderNO: %s, err: %v", orderNO, err)
				return "", i.T(ctx, errcode.CodeQueryOrderFailed)
			}

			// 构造资源对象
			resource := types.DecryptedResource{
				OutTradeNo:     orderNO,
				TransactionID:  *resp.TransactionId,
				TradeState:     *resp.TradeState,
				TradeStateDesc: *resp.TradeStateDesc,
				SuccessTime:    *resp.SuccessTime,
				Payer: struct {
					OpenID string `json:"openid"`
				}{
					OpenID: *resp.Payer.Openid,
				},
				Amount: struct {
					Total    int64  `json:"total"`
					Currency string `json:"currency"`
				}{
					Total:    *resp.Amount.Total,
					Currency: *resp.Amount.Currency,
				},
			}

			// 处理支付成功
			handlePaySuccess(ctx, orderInfo, resource)
			return string(consts.OrderStatusPayed), nil
		} else {
			glog.Warningf(ctx, "order NO (%s) not found in cache, but paid in wechat", orderNO)
			return "", i.T(ctx, errcode.CodeOrderNotFound)
		}
	case "NOTPAY", "USERPAYING": // 未支付或支付中
		return string(consts.OrderStatusNotPay), nil
	default: // 其他状态视为失败
		return string(consts.OrderStatusFailed), nil
	}
}

// 查询微信支付订单状态
func queryWeChatOrder(ctx context.Context, orderNO string) (string, error) {
	resp, _, err := nativeSvc.QueryOrderByOutTradeNo(ctx, native.QueryOrderByOutTradeNoRequest{
		OutTradeNo: core.String(orderNO),
		Mchid:      core.String(config.WeChatMchID),
	})
	if err != nil {
		return "", i.T(ctx, errcode.CodeQueryOrderFailed)
	}

	glog.Infof(ctx, "query wechat order success, orderNO: %s, tradeState: %s", orderNO, *resp.TradeState)
	return *resp.TradeState, nil
}

// 计算商品价格
func calculateProductPrice(ctx context.Context, productID string, productValue int) (int64, error) {
	var price int64 = 1
	switch consts.ProductID(productID) {
	case consts.ProductPlugin20:
		price = 2000 // 单位：分，下同
	case consts.ProductPlugin50:
		price = 5000
	case consts.ProductPlugin100:
		price = 10000
	case consts.ProductPlugin200:
		price = 20000
	case consts.ProductPlugin500:
		price = 50000
	case consts.ProductPlugin1000:
		price = 100000
	case consts.ProductCustom:
		price = int64(productValue) * 100 // 单位：分，WebUI传入的单位是：元
	case consts.ProductVipYear: // VIP年卡，199元
		price = 19900 // 199元 = 19900分
	case consts.ProductSVipMonth: // SVIP月卡，59元
		price = 5900 // 59元 = 5900分
	case consts.ProductSVipYear: // SVIP年卡，398元
		price = 39800 // 398元 = 39800分
	default:
		return 0, i.T(ctx, errcode.CodeInvalidProduct)
	}
	return price, nil
}

// 根据充值金额计算赠送用户积分，返回实际赠送积分
func calcRechargeWithBonus(amount int64) float64 {
	// 充值金额，单位为分，例如5000表示50元

	if amount >= 100000 { // 充值1000元以上，赠送200积分
		return 200
	}

	if amount >= 50000 { // 充值500元以上，赠送75积分
		return 75
	}

	if amount >= 20000 { // 充值200元以上，赠送20积分
		return 20
	}

	if amount >= 10000 { // 充值100元以上，赠送8积分
		return 8
	}

	if amount >= 5000 { // 充值50元以上，赠送3积分
		return 3
	}

	return 0 // 其它充值，不赠送积分
}

// 处理支付成功逻辑（公共函数）
func handlePaySuccess(ctx context.Context, orderInfo types.OrderInfo, resource types.DecryptedResource) {
	orderNO := resource.OutTradeNo
	key := fmt.Sprintf(consts.OrderNOKeyFormat, orderNO)

	// 1. 验证订单金额
	if orderInfo.Price != resource.Amount.Total {
		glog.Warningf(ctx, "order price not match, userID: %d, orderNO: %s, transactionID: %s, price: %d, total: %d",
			orderInfo.UserID, orderNO, resource.TransactionID, orderInfo.Price, resource.Amount.Total)
		// 金额不匹配，直接返回，避免错误充值
		return
	}

	// 2. 计算赠送积分
	bonus := calcRechargeWithBonus(resource.Amount.Total)
	rechargeAmount := float64(resource.Amount.Total) / 100.0 // 充值金额（元）
	totalPoints := rechargeAmount + bonus                    // 总积分 = 充值金额 + 赠送积分

	// 3. 使用事务处理，确保数据一致性
	tx := db.NewDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			glog.Errorf(ctx, "handlePaySuccess panic: %v, orderNO: %s", r, orderNO)
		}
	}()

	// 4. 创建订单记录（使用事务）
	order := model.Order{
		UserID:        orderInfo.UserID,
		ProductID:     string(orderInfo.ProductID),
		OrderNO:       orderNO,
		PayMethod:     string(consts.PayMethodWeChat),
		Amount:        rechargeAmount, // 这里记录真实的订单金额，不包含赠送积分
		TransactionID: resource.TransactionID,
		PayedAt:       time.Now(),
		OrderStatus:   string(consts.OrderStatusPayed),
		Payer:         resource.Payer.OpenID,
	}
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()

		// 判断是否为唯一索引冲突 - 如果冲突，说明已经处理过，不需要重复处理
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			glog.Infof(ctx, "order already processed, userID:%d, orderNO:%s, transactionID:%s, price: %.2f, err: %+v",
				orderInfo.UserID, orderNO, resource.TransactionID, rechargeAmount, err)
		} else {
			glog.Errorf(ctx, "create order failed, userID:%d, orderNO:%s, transactionID:%s, price: %.2f, err: %v",
				orderInfo.UserID, orderNO, resource.TransactionID, rechargeAmount, err)
		}

		return
	}

	// 5. 根据产品类型处理不同的逻辑
	err := processPluginRecharge(ctx, tx, orderInfo, orderNO, rechargeAmount, bonus, totalPoints)
	if err != nil {
		tx.Rollback()
		return
	}

	// 6. 提交事务
	if err := tx.Commit().Error; err != nil {
		glog.Errorf(ctx, "commit transaction failed, orderNO: %s, err: %v", orderNO, err)
		return
	}

	// 7. 更新缓存状态（事务成功后才更新）
	orderInfo.OrderStatus = consts.OrderStatusPayed
	cache.GetInstance().Set(key, orderInfo, consts.PayTimeout)

	// 8. 发送支付成功邮件（异步操作，失败不影响主流程）
	helper.SendPaymentSuccessEmail(ctx, orderInfo.UserID, resource.OutTradeNo, rechargeAmount, resource.TransactionID, resource.SuccessTime)
	glog.Infof(ctx, "handle pay success completed, orderNO: %s, rechargeAmount: %.2f, bonus: %.2f, totalPoints: %.2f",
		orderNO, rechargeAmount, bonus, totalPoints)
}

// 处理插件充值逻辑
func processPluginRecharge(ctx context.Context, tx *gorm.DB, orderInfo types.OrderInfo, orderNO string, rechargeAmount, bonus, totalPoints float64) error {
	// 根据产品类型决定如何处理
	switch orderInfo.ProductID {
	case consts.ProductVipYear, consts.ProductSVipMonth, consts.ProductSVipYear:
		// 处理VIP产品
		return processVIPUpgrade(ctx, tx, orderInfo, orderNO, rechargeAmount)
	default:
		// 处理普通插件充值
		return processNormalRecharge(ctx, tx, orderInfo, orderNO, rechargeAmount, bonus, totalPoints)
	}
}

// 处理普通插件充值逻辑
func processNormalRecharge(ctx context.Context, tx *gorm.DB, orderInfo types.OrderInfo, orderNO string, rechargeAmount, bonus, totalPoints float64) error {
	// 更新用户积分（使用原子操作，避免并发问题）
	result := tx.Model(&model.User{}).Where("id = ?", orderInfo.UserID).Update("points", gorm.Expr("points + ?", totalPoints))
	if result.Error != nil {
		glog.Errorf(ctx, "update user points failed, userID: %d, orderNO: %s, totalPoints: %.2f, err: %v",
			orderInfo.UserID, orderNO, totalPoints, result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		glog.Errorf(ctx, "user not found, userID: %d, orderNO: %s", orderInfo.UserID, orderNO)
		return errors.New("user not found")
	}

	// 记录充值日志（分两条记录：充值本金 + 赠送积分）
	// 1. 记录充值本金
	blRecharge := &model.PluginBalanceLog{
		UserID:      orderInfo.UserID,
		OrderNO:     orderNO,
		Points:      rechargeAmount,
		Description: consts.PluginRechargeDesc,
	}
	if err := tx.Create(blRecharge).Error; err != nil {
		glog.Errorf(ctx, "create recharge balance log failed, orderInfo: %+v, err: %v", orderInfo, err)
		return err
	}

	// 2. 如果有赠送，记录赠送日志
	if bonus > 0 {
		blBonus := &model.PluginBalanceLog{
			UserID:      orderInfo.UserID,
			OrderNO:     orderNO,
			Points:      bonus,
			Description: consts.PluginRechargeBonusDesc,
		}
		if err := tx.Create(blBonus).Error; err != nil {
			glog.Errorf(ctx, "create bonus balance log failed, orderInfo: %+v, err: %v", orderInfo, err)
			return err
		}
	}
	return nil
}

// 处理VIP升级逻辑
func processVIPUpgrade(ctx context.Context, tx *gorm.DB, orderInfo types.OrderInfo, orderNO string, rechargeAmount float64) error {
	// 获取当前用户信息
	var user model.User
	if err := tx.Where("id = ?", orderInfo.UserID).First(&user).Error; err != nil {
		glog.Errorf(ctx, "get user failed, userID: %d, orderNO: %s, err: %v", orderInfo.UserID, orderNO, err)
		return err
	}

	// 确定VIP等级和续费时长
	var vipLevel int
	var duration time.Duration
	var desc string
	switch orderInfo.ProductID {
	case consts.ProductVipYear:
		vipLevel = 1                    // VIP_YEAR
		duration = 365 * 24 * time.Hour // 1年
		desc = "升级VIP年卡会员"
	case consts.ProductSVipMonth:
		vipLevel = 2                   // SVIP_MONTH
		duration = 30 * 24 * time.Hour // 1个月
		desc = "升级SVIP月卡会员"
	case consts.ProductSVipYear:
		vipLevel = 3                    // SVIP_YEAR
		duration = 365 * 24 * time.Hour // 1年
		desc = "升级SVIP年卡会员"
	default:
		glog.Errorf(ctx, "unknown VIP product, productID: %s", orderInfo.ProductID)
		return errors.New("unknown VIP product")
	}

	// 计算新的VIP到期时间
	var newExpireAt time.Time
	currentTime := time.Now()
	if user.VIPExpireAt.After(currentTime) {
		// 如果当前VIP还未过期，则在当前基础上延长
		newExpireAt = user.VIPExpireAt.Add(duration)
	} else {
		// 如果当前VIP已过期或未开通，则从现在开始计算
		newExpireAt = currentTime.Add(duration)
	}

	// 更新用户VIP信息
	updates := map[string]interface{}{
		"vip_level":     vipLevel,
		"vip_expire_at": newExpireAt,
	}
	result := tx.Model(&model.User{}).Where("id = ?", orderInfo.UserID).Updates(updates)
	if result.Error != nil {
		glog.Errorf(ctx, "update user VIP info failed, userID: %d, orderNO: %s, err: %v",
			orderInfo.UserID, orderNO, result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		glog.Errorf(ctx, "user not found when updating VIP info, userID: %d, orderNO: %s", orderInfo.UserID, orderNO)
		return errors.New("user not found")
	}

	// 记录VIP升级日志
	blVIPUpgrade := &model.PluginBalanceLog{
		UserID:      orderInfo.UserID,
		OrderNO:     orderNO,
		Points:      rechargeAmount,
		Description: desc,
	}
	if err := tx.Create(blVIPUpgrade).Error; err != nil {
		glog.Errorf(ctx, "create VIP upgrade balance log failed, orderInfo: %+v, err: %v", orderInfo, err)
		return err
	}

	return nil
}

// 解密微信回调资源数据
func decryptWeChatResource(apiKey, associatedData, nonce []byte, ciphertext string) ([]byte, error) {
	// 检查密钥长度 (32字节)
	if len(apiKey) != 32 {
		return nil, gerror.Newf("invalid APIv3 key length: expected 32, got %d", len(apiKey))
	}

	// Base64解码密文
	cipherBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed: %w", err)
	}

	// 创建AES-GCM解密器
	block, err := aes.NewCipher(apiKey)
	if err != nil {
		return nil, fmt.Errorf("create cipher failed: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("create GCM failed: %w", err)
	}

	// 执行解密 (GCM标准需要12字节nonce)
	if len(nonce) != gcm.NonceSize() {
		return nil, fmt.Errorf("invalid nonce size: expected %d, got %d",
			gcm.NonceSize(), len(nonce))
	}

	plaintext, err := gcm.Open(nil, nonce, cipherBytes, associatedData)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}

	return plaintext, nil
}
