package internal

import (
	"path"
	"runtime"
	"strings"
)

/*
	use: Copyright(make([]uintptr, 1))
	return: main.xxx.xxx.xxx
*/
func Copyright(pc []uintptr) string {
	s := ""
	for i := 1; i < 4; i++ {
		runtime.Callers(i, pc)
		if i == 3 {
			s += strings.Replace(path.Ext(runtime.FuncForPC(pc[0]).Name()), ".", "", -1)
		} else {
			s += strings.Replace(path.Ext(runtime.FuncForPC(pc[0]).Name()), ".", "", -1) + " 》"
		}
	}
	return s
}
