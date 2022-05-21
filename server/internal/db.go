package internal

import (
	. "github.com/Ran-Xing/ATS/server/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() {

	if DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       viper.GetString("MYSQL.HOST"), // DSN data source name
		DefaultStringSize:         256,                           // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                          // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                          // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                          // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                         // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		// TODO Debug SQL Model
		//Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		log.Errorf("Connect DataBase Error: %v", err)
	}
	autoMigrate(&UserInfo{})
	//autoMigrate(&Hotel{})
	//autoMigrate(&Room{})
	//autoMigrate(&User{})
	//autoMigrate(&Booking{})
}
func autoMigrate(table interface{}) {
	if err = DB.AutoMigrate(table); err != nil {
		log.Errorf("AutoMigrate Error: %v", err)
		return
	}
}
