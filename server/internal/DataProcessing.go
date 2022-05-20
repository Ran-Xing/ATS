package internal

import (
	log "github.com/sirupsen/logrus"
	"os"
)

/*
	FileInit 初始化文件/夹
*/
func FileInit(fileAttributes bool, fileName string) {
	if _, err = os.Stat(fileName); err != nil {
		if fileAttributes {
			if err = os.MkdirAll(fileName, os.ModePerm); err != nil {
				log.Errorf("[FileInit] 创建[%v] 目录失败, Error: [%v] CoptRight: [%s]", fileName, err, Copyright(make([]uintptr, 1)))
			} else {
				log.Infof("[FileInit] 创建[%v] 目录成功", fileName)
			}
		} else {
			var f *os.File
			if f, err = os.Create(fileName); err != nil {
				log.Errorf("[FileInit] 创建[%v] 文件失败, Error: [%v] CoptRight: [%s]", fileName, err, Copyright(make([]uintptr, 1)))
			} else {
				log.Infof("[FileInit] 创建[%v] 文件成功", fileName)
			}
			defer func(f *os.File) {
				if err = f.Close(); err != nil {
					log.Errorf("[FileInit] 关闭[%v] 文件失败, Error: [%v] CoptRight: [%s]", fileName, err, Copyright(make([]uintptr, 1)))
				}
			}(f)
		}
	}
}
