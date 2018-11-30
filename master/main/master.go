package main

import (
	"runtime"

	"fmt"

	"flag"

	"bx.go/learngo/cron/crontab/master"
)

var (
	confFile string
)

func initArgs() {
	flag.StringVar(&confFile, "config", "./master.json", "传入配置文件路径")
	flag.Parse()
}

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)

	initArgs()

	//初始化线程
	initEnv()

	//加载配置
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}

	//任务管理器
	if err = master.InitJobMgr(); err != nil {
		goto ERR
	}

	//启动Api HTTP服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	//正常退出
	return

ERR:
	fmt.Println(err)
}
