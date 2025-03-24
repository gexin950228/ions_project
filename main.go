package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	_ "ions_project/routers"
	"ions_project/utils"
)

func main() {
	orm.RunCommand()

	// 开启session
	beego.BConfig.WebConfig.Session.SessionOn = true

	// 开启未登录请求拦截
	beego.InsertFilter("/main/*", beego.BeforeRouter, utils.LoginFilter)

	// 开启日志
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"logs/app.log", "separate":["error", "info", "panic"]}`)
	if err != nil {
		logs.Error(err)
	}
	logs.SetLogFuncCallDepth(3)
	beego.SetStaticPath("/upload", "upload")

	// 开启sql打印
	orm.Debug = true
	beego.Run(":8089")
}
