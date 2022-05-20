package model

import (
	"gorm.io/gorm"
	"time"
)

type (
	LoginInfo struct {
		UserName string `json:"username" binding:"required" gorm:"column:username;comment:用户名;primaryKey"` // 用户名称
		PassWord string `json:"password" binding:"required" gorm:"column:password;comment:用户密码"`           // 用户密码
		Level    int64  `json:"level" default:"0" gorm:"column:level;comment:用户等级"`                        // 用户等级
	}
	UserInfo struct {
		// import LoginInfo exclude PassWord
		LoginInfo
		CreatedAt time.Time      `json:"CreatedAt" gorm:"column:created_at;comment:创建时间"` // 创建时间
		UpdatedAt time.Time      `json:"UpdatedAt" gorm:"column:updated_at;comment:更新时间"` // 更新时间
		Deleted   gorm.DeletedAt `json:"Deleted" gorm:"column:deleted;comment:删除时间"`      // 删除时间
	}
)
