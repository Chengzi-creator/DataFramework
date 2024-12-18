package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 定义密钥（保存在安全位置）
var jwtKey = []byte("my_secret_key")

// Claims 定义 JWT 的 Payload 结构
type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT Token
func GenerateToken(userID int, username string) (string, error) {
	// 设置 Token 过期时间
	expirationTime := time.Now().Add(24 * time.Hour)

	// 创建 JWT Claims，包含用户信息和过期时间
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// 使用 HS256 签名算法生成 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateToken 验证 JWT Token
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// 解析 Token 并验证签名
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
