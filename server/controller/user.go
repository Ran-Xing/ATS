package controller

import (
	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	. "grs/internal"
	. "grs/models"
	. "grs/response"
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
	if ok, err := regexp2.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`, 0).MatchString(l.Email); !ok || err != nil || len(l.Email) > 50 {
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

	// 设置默认用户名
	if match, err := regexp2.MustCompile(`^\w+([-+.]\w+)*(?=@)`, 0).FindStringMatch(l.Email); err != nil || match == nil {
		log.Errorf("match email error, %s", err)
	} else {
		for match != nil {
			u.UserName = match.String()
			break
		}
	}

	// 数据复制 存储
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

	c.Set("UserInfo", u)

	// 清除不需要的数据
	u.PassWord = ""
	Success(c, gin.H{"message": "获取用户信息成功!", "token": token, "UserInfo": u})
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
	if ok, err := regexp2.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`, 0).MatchString(l.Email); !ok || err != nil || len(l.Email) > 50 {
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

	// 清除不需要的数据
	u.PassWord = ""
	Success(c, gin.H{"message": "获取用户信息成功!", "token": token, "UserInfo": u})
}

func Info(c *gin.Context) {
	var (
		u  UserInfo
		ok bool
	)
	var ut interface{}
	if ut, ok = c.Get("UserInfo"); !ok || ut == nil {
		Fail(c, "获取用户信息失败!")
		return
	}
	u = ut.(UserInfo)
	u.PassWord = ""
	Success(c, gin.H{"message": "获取用户信息成功!", "UserInfo": u})
}

func UpdateProfile(c *gin.Context) {
	var (
		u   UserInfo
		l   UpdateUserInfo
		ok  bool
		num int64
		err error
	)
	var ut interface{}
	if ut, ok = c.Get("UserInfo"); !ok || ut == nil {
		Fail(c, "获取用户信息失败!")
		return
	}
	//u = ut.(UserInfo)
	//u.PassWord = ""

	// 获取需要修改的参数
	if err = c.ShouldBindJSON(&l); err != nil {
		Fail(c, "参数错误")
		return
	}

	// 用户想搞空白名？
	if u.UserName == "" {
		u.UserName = ut.(UserInfo).UserName
	}

	// 匹配 电子邮箱
	if l.Email != "" {
		if ok, err := regexp2.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`, 0).MatchString(l.Email); !ok || err != nil || len(l.Email) > 50 {
			Fail(c, "邮箱格式错误")
			return
		}
	} else {
		u.Email = ut.(UserInfo).Email
	}

	// 匹配密码
	if l.PassWord != "" {
		if num = int64(len(l.PassWord)); num < 6 || num > 30 {
			Fail(c, "密码长度不能小于6位")
			return
		}
	} else {
		u.PassWord = ut.(UserInfo).PassWord
	}

	//if num = int64(len(l.OldPassWord)); num < 6 || num > 30 {
	//	Fail(c, "密码长度不能小于6位")
	//	return
	//}

	// 匹配手机号
	if l.OldPhone != "" {
		if ok, err := regexp2.MustCompile("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$", 0).MatchString(l.Email); !ok || err != nil || len(l.Email) > 50 {
			Fail(c, "手机格式错误")
			return
		}
	} else {
		u.Phone = ut.(UserInfo).Phone
	}

	// 查询用户
	if err = DB.Model(&UserInfo{}).Where("Email = ?", l.Email).Count(&num).Error; err != nil {
		log.Errorf("查询失败, %s", err)
		Fail(c, "查询失败")
		return
	}
	if num == 0 {
		Fail(c, "用户不存在")
		return
	}

	//// 查询密码 TODO 密码验证
	//if err = bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(l.OldPassWord)); err != nil {
	//	Fail(c, "密码错误")
	//	return
	//}
	if err = DB.Model(&UserInfo{}).Where("Email = ?", l.Email).Updates(&u).Error; err != nil {
		log.Errorf("查询失败, %s", err)
		Fail(c, "更新失败")
		return
	}

	// 创建备份数据库
	l.Email, l.OldEmail = updateProfileDB(u.Email, l.Email)
	l.PassWord, l.OldPassWord = updateProfileDB(u.PassWord, l.PassWord)
	l.UserName, l.OldUserName = updateProfileDB(u.UserName, l.UserName)
	l.Phone, l.OldPhone = updateProfileDB(u.Phone, l.Phone)

	// 查询用户旧信息
	num = 0
	if err = DB.Model(&UpdateUserInfo{}).Where("Email = ?", l.Email).Count(&num).Error; err != nil {
		log.Errorf("查询失败, error: %s", err)
		return
	}

	// 创建用户旧信息表
	if num == 0 {
		if err = DB.Create(&l).Error; err != nil {
			log.Errorf("创建用户旧信息表失败, error: %v", err)
			return
		}
	}

	// 更新用户旧表信息

	if err = DB.Model(&UpdateUserInfo{}).Where("Email = ?", l.Email).Updates(&l).Error; err != nil {
		log.Errorf("更新用户旧表信息失败, error: %v", err)
	}
}

/*
	updateProfileDB(old,new) (new,old)
*/
func updateProfileDB(u, l string) (string, string) {
	var v string
	if l != "" {
		v += u + "|" + l
		return l, v
	}
	return u, ""
}
