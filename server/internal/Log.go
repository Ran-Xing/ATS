package internal

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func LogInit() {
	log.SetPrefix("\x1b[1;32m[wechatBot] \x1b[0m")
	// \x1b[%dm%s\x1b[0m
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	FileInit(true, viper.GetString("LogPath"))
	logrus.SetLevel(logrus.WarnLevel | logrus.InfoLevel | logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2006-01-02 15:04:05",
		FullTimestamp:             true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := fmt.Sprintf("%v:%v", path.Base(frame.File), strconv.Itoa(frame.Line))
			function = fmt.Sprintf(" [%v]", strings.Replace(path.Ext(frame.Function), ".", "", -1))
			return fileName, function
		},
	})
}
