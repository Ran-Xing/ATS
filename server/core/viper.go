package core

import (
	"fmt"
	"github.com/Ran-Xing/ATS/server/internal"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	var (
		err error
		v   = viper.New()
	)
	v.SetConfigFile(viper.ConfigFileUsed())
	if err = viper.ReadInConfig(); err != nil {
		log.Errorf("[Viper] 读取配置文件失败, Error: [%v] CoptRight: [%s]", err, internal.Copyright(make([]uintptr, 1)))
	} else {
		log.Infof("[Viper] 读取配置文件成功")
	}
	fmt.Println(v.AllSettings())
	fmt.Println(viper.AllSettings())
	// 监控配置和重新获取配置
	// var err error
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("[Viper] Config Changed")
	})
	return v
}
