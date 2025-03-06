package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "ions_project/routers"
)

func init() {
	username := beego.AppConfig.String("username")
	pwd := beego.AppConfig.String("pwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	db := beego.AppConfig.String("db")

	// username:pwd@tcp(ip:port)/db?charset=utf8&loc=Local
	dataSource := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8&loc=Local"
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Println(err)
	}
	err = orm.RegisterDataBase("default", "mysql", dataSource)
	if err != nil {
		fmt.Println(err)
	}

	ret := fmt.Sprintf("host:%s|port:%s|db:%s", host, port, db)
	logs.Info(ret)
	//orm.RegisterModel(
	//	new(auth.Auth),
	//	new(auth.Role),
	//	new(auth.User),
	//	new(auth.CarBrand),
	//	new(auth.Cars),
	//	new(auth.CarsApply),
	//	new(auth.MessageNotify),
	//)
}

func main() {
	orm.RunCommand()

	// 开启session
	beego.BConfig.WebConfig.Session.SessionOn = true

	// 开启未登录请求拦截
	//beego.InsertFilter("/main/*", beego.BeforeRouter, utils.LoginFilter)

	// 开启日志
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"logs/app.log", "separate":["error", "info", "panic"]}`)
	if err != nil {
		fmt.Println(err)
	}
	logs.SetLogFuncCallDepth(3)
	beego.SetStaticPath("/upload", "upload")

	// 开启sql打印
	orm.Debug = true

	beego.Run(":8089")
}
