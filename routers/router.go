package routers

import (
	"github.com/astaxie/beego"
	"ions_project/controllers"
	"ions_project/controllers/auth"
	"ions_project/controllers/cars"
	"ions_project/controllers/login"
	"ions_project/controllers/user"
)

func init() {
	// 不需要登录既可请求的url
	beego.Router("/", &login.LController{})
	beego.Router("/main/user/log_out", &login.LController{}, "get:LogOut")
	beego.Router("/change_captcha", &login.LController{}, "get:ChangeCaptcha")

	// 必须登录才可请求的url

	// 后台首页
	beego.Router("/main/index", &controllers.HomeController{})
	beego.Router("/main/index/notify", &controllers.HomeController{}, "get:NotifyList")
	beego.Router("/main/index/read_notify", &controllers.HomeController{}, "get:ReadNotify")
	beego.Router("/main/welcome", &controllers.HomeController{}, "get:Welcome")

	// user模块
	beego.Router("/main/user/list", &user.UserController{}, "get:List")
	beego.Router("/main/user/to_add", &user.UserController{}, "get:ToAdd")
	beego.Router("/main/user/do_add", &user.UserController{}, "post:DoAdd")
	beego.Router("/main/user/is_active", &user.UserController{}, "post:IsActive")
	beego.Router("/main/user/delete", &user.UserController{}, "get:Delete")
	beego.Router("/main/user/reset_pwd", &user.UserController{}, "get:ResetPassword")
	beego.Router("/main/user/to_edit", &user.UserController{}, "get:ToUpdate")
	beego.Router("/main/user/do_edit", &user.UserController{}, "post:DoUpdate")
	beego.Router("/main/user/muli_delete", &user.UserController{}, "post:MuliDelete")

	// auth模块
	beego.Router("/main/auth/list", &auth.AuthController{}, "get:List")
	beego.Router("/main/auth/to_auth_add", &auth.AuthController{}, "get:ToAuthAdd")
	beego.Router("/main/auth/auth_add", &auth.AuthController{}, "post:DoAdd")

	// cars
	beego.Router("/main/cars/list", &cars.CarsController{}, "get:List")

}
