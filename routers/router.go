package routers

import (
	"github.com/astaxie/beego"
	"ions_project/controllers"
	"ions_project/controllers/login"
)

func init() {
	// 不需要登录既可请求的url
	beego.Router("/", &login.LController{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/main/user/log_out", &login.LController{}, "get:LogOut")
	beego.Router("/change_captcha", &login.LController{}, "get:ChangeCaptcha")
}
