package middleware

import (
	"github.com/gin-gonic/gin"
	"go_mall/pkg/e"
	"go_mall/pkg/utils"
	"go_mall/serializer"
	"net/http"
	"time"
)

// JWT
// @Description: jwt中间件
// @return gin.HandlerFunc
func JWT() gin.HandlerFunc {
	// token验证中间件
	return func(c *gin.Context) {
		code := e.Success
		claims, err := utils.ParseJWT(c.GetHeader("Authorization"))
		if err != nil { // 解析失败,过期了也会解析失败
			code = e.ErrorWithParseToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorWithExpiredToken
		}
		if code != e.Success { // 不成功
			c.Abort()
			c.JSON(http.StatusUnauthorized,
				serializer.Response{
					Code:    code,
					Message: e.GetMessageByCode(code),
				})
			return
		}
		// 解析成功，将claims的值进行传递
		c.Set("claims", claims)
		c.Next()
	}
}
