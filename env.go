package common

import (
	"fmt"
	"log"
)

const (
	prodPath = "../"     // 生产环境路径
	devPath  = "../res/" // 开发环境路径
)

// 程序启动错误处理
func Error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 运行开发环境
func RunDev(isHttp bool, port string) {
	if isDev {
		logOut(isHttp, port)
	}
}

// 获取文件夹路径
func env(param string) string {
	if isDev {
		return devPath + param
	} else {
		return prodPath + param
	}
}

// 开发环境日志输出
func logOut(isHttp bool, port string) {
	if isHttp {
		fmt.Printf("Listening and serving HTTPS on %v\n", port)
	} else {
		fmt.Printf("Listening and serving HTTP on %v\n", port)
	}
}
