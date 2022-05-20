package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	. "grs/internal"
	. "grs/models"
	. "grs/response"
	"regexp"
)

func Register(c *gin.Context) {
	var (
		u   UserInfo
		l   LoginInfo
		num int64
		err error
	)
	if err = c.ShouldBindJSON(&l); err != nil {
		Fail(c, "参数错误")
		return
	}
	// TODO 默认值操作

	// 匹配 电子邮箱
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(l.Email) || len(l.Email) > 50 {
		Fail(c, "邮箱格式错误")
		return
	}

	// 匹配密码
	if num = int64(len(l.PassWord)); num < 6 || num > 30 {
		Fail(c, "密码长度不能小于6位")
		return
	}

	// 创建默认用户
	if l.Level != 0 {
		Fail(c, "用户等级错误")
		return
	}

	// 查询用户
	if err = DB.Model(&UserInfo{}).Where("Email = ?", l.Email).Count(&num).Error; err != nil {
		log.Errorf("查询失败, %s", err)
		Fail(c, "查询失败")
		return
	}

	if num > 0 {
		Fail(c, "用户已存在")
		return
	}

	// 创建用户
	var hashPass []byte
	if hashPass, err = bcrypt.GenerateFromPassword([]byte(l.PassWord), bcrypt.DefaultCost); err != nil {
		log.Errorf("hash password error, %s", err)
		Fail(c, "密码加密失败")
		return
	}
	l.PassWord = string(hashPass)

	// 数据复制
	u.LoginInfo = l
	if err = DB.Create(&u).Error; err != nil {
		Fail(c, "注册失败")
		return
	}

	// 生成token
	var token string
	if token, err = JwtEncryption(l.Email); err != nil {
		Fail(c, "系统异常")
		return
	}

	// 清除不需要的数据
	Success(c, gin.H{"message": "注册成功", "token": token})
}

func Login(c *gin.Context) {
	var (
		u   UserInfo
		l   LoginInfo
		num int64
		err error
	)
	if err = c.ShouldBindJSON(&l); err != nil {
		Fail(c, "参数错误")
		return
	}

	// 匹配 电子邮箱
	reg := regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	if !reg.MatchString(l.Email) || len(l.Email) > 50 {
		Fail(c, "邮箱格式错误")
		return
	}

	// 匹配密码
	if num = int64(len(l.PassWord)); num < 6 || num > 30 {
		Fail(c, "密码长度不能小于6位")
		return
	}

	// 查询用户
	if err = DB.Model(&UserInfo{}).Where("Email = ?", l.Email).Count(&num).Scan(&u).Error; err != nil {
		log.Errorf("查询失败, %s", err)
		Fail(c, "查询失败")
		return
	}
	if num == 0 {
		Fail(c, "用户不存在")
		return
	}

	// 查询密码
	if err = bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(l.PassWord)); err != nil {
		Fail(c, "密码错误")
		return
	}

	// 生成token
	var token string
	if token, err = JwtEncryption(l.Email); err != nil {
		Fail(c, "系统异常")
		return
	}

	c.Set("UserInfo", u)
	Success(c, gin.H{"message": "登录成功", "token": token, "Email": l.Email})
}

func Info(c *gin.Context) {
	var (
		u  UserInfo
		ok bool
	)
	var ut interface{}
	if ut, ok = c.Get("UserInfo"); !ok || ut == nil {
		Fail(c, "请登录!")
		return
	}
	u = ut.(UserInfo)
	u.PassWord = ""
	Success(c, gin.H{"message": "获取用户信息成功!", "UserInfo": u})
}
