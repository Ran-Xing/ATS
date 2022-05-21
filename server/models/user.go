package model

import (
	"gorm.io/gorm"
	"time"
)

type (
	LoginInfo struct {
		Email    string `json:"email" binding:"required" gorm:"column:email;comment:邮箱;primaryKey"` // 用户名称
		PassWord string `json:"password" binding:"required" gorm:"column:password;comment:用户密码"`    // 用户密码
		Level    int64  `json:"level" default:"0" gorm:"column:level;comment:用户等级"`                 // 用户等级
	}
	UserInfo struct {
		// import LoginInfo exclude PassWord
		LoginInfo
		UserName  string         `json:"username" binding:"required" gorm:"column:username;comment:用户名;primaryKey"` // 用户名称
		Phone     string         `json:"phone" binding:"required" gorm:"column:phone;comment:用户手机号"`                // 用户手机号
		Age       int64          `json:"age" gorm:"column:age;comment:年龄"`                                          // 年龄
		Education string         `json:"education" gorm:"column:education;comment:学历"`                              // 学历
		Blog      string         `json:"blog" gorm:"column:blog;comment:博客地址"`                                      // 博客地址
		Github    string         `json:"github" gorm:"column:github;comment:github地址"`                              // github地址
		CreatedAt time.Time      `json:"CreatedAt" gorm:"column:created_at;comment:创建时间"`                           // 创建时间
		UpdatedAt time.Time      `json:"UpdatedAt" gorm:"column:updated_at;comment:更新时间"`                           // 更新时间
		Deleted   gorm.DeletedAt `json:"Deleted" gorm:"column:deleted;comment:删除时间"`                                // 删除时间
	}
	UserResume struct {
		LoginInfo
		Resume string `json:"resume" gorm:"column:resume;comment:个人简历"` // 个人简历
	}
	UpdateUserInfo struct {
		UserInfo
		OldEmail    string `json:"oldEmail" gorm:"column:old_email;comment:旧邮箱"`         // 旧邮箱
		OldPassWord string `json:"old_password" gorm:"column:old_password;comment:旧密码"`  // 旧密码
		OldUserName string `json:"old_username" gorm:"column:old_username;comment:旧用户名"` // 旧用户名
		OldPhone    string `json:"old_phone" gorm:"column:old_phone;comment:旧手机号"`       // 旧手机号
	}
)
