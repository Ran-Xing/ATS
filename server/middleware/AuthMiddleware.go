package middleware

import (
	. "github.com/Ran-Xing/ATS/server/internal"
	. "github.com/Ran-Xing/ATS/server/models"
	. "github.com/Ran-Xing/ATS/server/response"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

var (
	u     UserInfo
	num   int64
	err   error
	email string
	ok    bool
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			Fail(c, "无权限访问")
			c.Abort()
			return
		}

		// 获取token
		if email, ok = JwtDecryption(tokenString[7:]); !ok {
			Fail(c, "无权限访问")
			c.Abort()
			return
		}

		// 查询用户信息
		if err = DB.Model(&UserInfo{}).Where("email = ?", email).Count(&num).Scan(&u).Error; err != nil {
			log.Errorf("查询失败, %s", err)
			Fail(c, "查询失败!")
			c.Abort()
			return
		}

		if num == 0 {
			Fail(c, "用户不存在!")
			c.Abort()
			return
		}

		// 查询用户权限
		if u.Email != email {
			Fail(c, "用户资料异常!")
			c.Abort()
			return
		}

		// 将用户存入上下文
		c.Set("UserInfo", u)
		c.Next()
	}
}
