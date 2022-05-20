package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	. "grs/internal"
	"grs/router"
)

var (
	err error
)

func init() {
	ViperInit()
	LogInit()
	ConnectDataBase()
}

func main() {
	r := Router.CollectRoute(gin.Default())
	if err = r.Run(":8081"); err != nil {
		log.Errorf("start server error: %v", err)
		return
	}
	if err = r.Run(viper.GetString("PORT")); err != nil {
		log.Errorf("Start Gin Error!")
		return
	}
}
