package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	. "grs/internal"
	. "grs/models"
	. "grs/response"
	"strings"
)

var (
	u        UserInfo
	num      int64
	err      error
	userName string
	ok       bool
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			//c.JSON(401, gin.H{"message": "无权限访问"})
			Fail(c, "无权限访问")
			c.Abort()
			return
		}

		// 获取token
		if userName, ok = JwtDecryption(tokenString[7:]); !ok {
			Fail(c, "无权限访问")
			c.Abort()
			return
		}

		// 查询用户信息
		if err = DB.Model(&UserInfo{}).Where("username = ?", userName).Count(&num).Scan(&u).Error; err != nil {
			log.Errorf("查询失败, %s", err)
			Fail(c, "查询失败")
			c.Abort()
			return
		}

		if num == 0 {
			Fail(c, "用户不存在")
			c.Abort()
			return
		}

		// 查询用户权限
		if u.UserName != userName {
			Fail(c, "用户不存在")
			c.Abort()
			return
		}

		// 将用户存入上下文
		c.Set("UserInfo", u)
		c.Next()
	}
}