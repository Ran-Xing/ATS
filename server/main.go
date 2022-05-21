package main

import (
	. "github.com/Ran-Xing/ATS/server/internal"
	. "github.com/Ran-Xing/ATS/server/router"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	r := CollectRoute(gin.Default())
	if err = r.Run(":8081"); err != nil {
		log.Errorf("start server error: %v", err)
		return
	}
	if err = r.Run(viper.GetString("PORT")); err != nil {
		log.Errorf("Start Gin Error!")
		return
	}
}
