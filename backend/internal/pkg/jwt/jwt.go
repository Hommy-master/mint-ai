package jwt

import (
	"context"
	"errors"
	"strconv"
	"time"

	"cozeos/internal/config"

	"github.com/gogf/gf/v2/os/glog"
	jwt "github.com/golang-jwt/jwt/v5"
)

// JWT claims结构
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// JWT服务结构
type JWTService struct {
	SigningKey []byte
	ExpireDays int
}

// 全局JWT服务实例
var jwtService *JWTService

// 初始化JWT服务
func init() {
	// 从环境变量中获取密钥，如果环境变量中没有则panic
	signingKey := config.JWTSigningKey
	if signingKey == "" {
		panic("JWT signing key is required but not provided")
	}

	// 从环境变量中获取过期天数，默认30天
	var expireDays int
	if config.JWTExpireDays != "" {
		if days, err := strconv.Atoi(config.JWTExpireDays); err == nil && days > 0 {
			expireDays = days
		} else {
			glog.Warningf(context.Background(), "Invalid JWT expire days: %s, using default 30 days", config.JWTExpireDays)
			expireDays = 30
		}
	} else {
		expireDays = 30
	}

	jwtService = &JWTService{
		SigningKey: []byte(signingKey),
		ExpireDays: expireDays,
	}
}

// 生成JWT令牌
func GenerateToken(userID uint, phone string, weChatID string) (string, error) {
	glog.Debugf(context.Background(), "Generating JWT token for user ID: %d, phone: %s, weChatID: %s", userID, phone, weChatID)

	// 设置令牌过期时间
	expireTime := time.Now().Add(time.Hour * 24 * time.Duration(jwtService.ExpireDays)) // 过期时间

	// 创建声明
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime), // 过期时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			Issuer:    "cozeos",                       // 签发者
			Subject:   "user_token",                   // 主题
		},
	}

	// 创建令牌对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成签名后的令牌
	tokenString, err := token.SignedString(jwtService.SigningKey)
	if err != nil {
		glog.Errorf(context.Background(), "Failed to generate JWT token: %v", err)
		return "", err
	}

	glog.Debugf(context.Background(), "JWT token generated successfully for user ID: %d", userID)
	return tokenString, nil
}

// 验证JWT令牌
func ValidateToken(tokenString string) (*Claims, error) {
	glog.Debugf(context.Background(), "Validating JWT token")

	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			glog.Errorf(context.Background(), "Unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New("unexpected signing method")
		}
		return jwtService.SigningKey, nil
	})

	if err != nil {
		glog.Errorf(context.Background(), "Failed to parse JWT token: %v", err)
		return nil, err
	}

	// 验证令牌是否有效
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		glog.Debugf(context.Background(), "JWT token validated successfully for user ID: %d", claims.UserID)
		return claims, nil
	}

	glog.Errorf(context.Background(), "Invalid JWT token")
	return nil, errors.New("invalid token")
}

// 从HTTP请求头中提取JWT令牌
func ExtractTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("authorization header is empty")
	}

	// 检查是否以"Bearer "开头
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		return authHeader[7:], nil
	}

	return "", errors.New("authorization header format must be Bearer {token}")
}
