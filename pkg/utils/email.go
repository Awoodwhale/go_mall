package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const EmailJWTExpireTime = 5 * time.Minute // 邮箱验证限时5分钟

// EmailClaims
// @Description: email的claims
type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

// GenerateEmailToken
// @Description: 生成email jwt
// @param uid uint
// @param operation uint
// @param email string
// @param password string
// @return string
// @return error
func GenerateEmailToken(uid, operation uint, email, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(EmailJWTExpireTime) // 5分钟的过期时间
	claims := EmailClaims{
		UserID:        uid,
		Email:         email,
		Password:      password,
		OperationType: operation,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Awoodwhale_email",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

// ParseEmailToken
// @Description: 解析email的jwt
// @param token string
// @return *EmailClaims
// @return error
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			// 只有token有效才返回claims
			return claims, nil
		}
	}
	return nil, err
}
