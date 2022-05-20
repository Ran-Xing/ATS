package internal

import (
	"bytes"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func ViperInit() {
	// 出厂设置
	RootPath, _ := os.Getwd()
	viper.Set("RootPath", RootPath)
	viper.Set("LogPath", RootPath+"/logs")
	// 初始化配置文件
	log.Printf("[Viper] Config Path: [%v]", viper.GetString("RootPath")+"/config.yaml")
	log.Printf("[Viper] Log Path: [%v]/", viper.GetString("LogPath"))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(viper.GetString("RootPath") + "/config")
	FileInit(false, viper.GetString("RootPath")+"/config/config.yaml")
	viperRead()
	if viper.GetString("PORT") == "" && viper.GetString("jwtSecret") == "" {
		var yamlExample = []byte(`
			jwtSecret: 02ace598-60d8-4634-9154-558d6018a288%
			PORT: 8081
			MYSQL:
				HOST: user:name@tcp(ip:port)/database_name?charset=utf8&parseTime=True&loc=Local
		`)
		if err = viper.ReadConfig(bytes.NewBuffer(yamlExample)); err != nil {
			log.Errorf("读取示例配置文件失败, Error: [%v] CoptRight: [%s]", err, Copyright(make([]uintptr, 1)))
			return
		}
		log.Infof("请检查配置文件! (CTRL + Z 撤销覆盖)")
	}

	viperWrite()
	// TODO 文件动态加载这个我也不会呀，懒得写了
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("[Viper] Config Changed")
		viperRead()
	})
}

/*
	ViperWrite()
	写入配置文件
*/
func viperWrite() {
	if err = viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
		log.Errorf("Write file Error: %v CoptRight: [%s]", err, Copyright(make([]uintptr, 1)))
	} else {
		log.Infof("Write file Success")
	}
}

func viperRead() {
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Errorf("配置文件丢失!")
		} else {
			log.Errorf("读取配置文件失败, Error: [%v] CoptRight: [%s]", err, Copyright(make([]uintptr, 1)))
		}
	} else {
		log.Infof("读取配置文件成功")
	}
}
