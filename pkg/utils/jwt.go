package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("woodwhale&sheepbotany")

const TokenExpireTime = 24 * time.Hour

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string, authority int) (string, error) {
	/**
	 * GenerateToken
	 * @Description: 签发Token
	 * @param id
	 * @param username
	 * @param authority
	 * @return string
	 * @return error
	 */
	nowTime := time.Now()
	expireTime := nowTime.Add(TokenExpireTime) // 24小时的过期时间
	claims := Claims{
		ID:        id,
		UserName:  username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Awoodwhale",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

func ParseToken(token string) (*Claims, error) {
	/**
	 * ParseToken
	 * @Description: 验证token
	 * @param token
	 * @return *Claims
	 * @return error
	 */
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			// 只有token有效才返回claims
			return claims, nil
		}
	}
	return nil, err
}