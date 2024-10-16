package jwt

import (
	"time"

	"github.com/PandaPy/pginer/template/initialize/config"
	"github.com/PandaPy/pginer/template/models"
	"github.com/golang-jwt/jwt/v5"
)

type JWTTokenResult struct {
	JWTToken  string `json:"jwtToken"`   // JWT 令牌
	ExpiresAt int64  `json:"expires_at"` // 令牌过期时间（Unix 时间戳）
	IssuedAt  int64  `json:"issued_at"`  // 令牌签发时间（Unix 时间戳）
}

func GenerateJWTToken(model *models.UserModel) (JWTTokenResult, error) {
	// 设置到期时间和签发时间为 Unix 时间戳
	expiresAtUnix := time.Now().Add(time.Hour * 72).Unix()
	issuedAtUnix := time.Now().Unix()

	// 定义 JWT 的 claims
	claims := jwt.MapClaims{
		"id":  model.ID,
		"exp": expiresAtUnix, // 到期时间戳
		"iat": issuedAtUnix,  // 签发时间戳
	}

	// 创建并签名 JWT 令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfig.Server.SecretKey))
	if err != nil {
		return JWTTokenResult{}, err
	}

	// 返回 JWTTokenResult 结构体
	return JWTTokenResult{
		JWTToken:  tokenString,
		ExpiresAt: expiresAtUnix,
		IssuedAt:  issuedAtUnix,
	}, nil
}
