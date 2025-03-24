package routers

import (
	"github.com/astaxie/beego"
	"ions_project/controllers"
	"ions_project/controllers/auth"
	"ions_project/controllers/caiwu"
	"ions_project/controllers/cars"
	"ions_project/controllers/login"
	"ions_project/controllers/news"
	"ions_project/controllers/user"
)

func init() {
	// 不需要登录既可请求的url
	beego.Router("/", &login.LController{})
	beego.Router("/main/user/log_out", &login.LController{}, "get:LogOut")
	beego.Router("/change_captcha", &login.LController{}, "get:ChangeCaptcha")

	// 必须登录才可请求的url

	// 后台首页
	beego.Router("/main/index/", &controllers.HomeController{})
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
	beego.Router("/main/auth/role_list", &auth.RoleController{}, "get:List")
	beego.Router("/main/auth/to_add", &auth.RoleController{}, "get:ToAdd")

	// 角色模块
	beego.Router("/main/role/list", &auth.RoleController{}, "get:List")
	beego.Router("/main/role/to_add", &auth.RoleController{}, "get:ToAdd")
	beego.Router("/main/role/do_add", &auth.RoleController{}, "post:DoAdd")
	beego.Router("/main/role/is_active", &auth.RoleController{}, "post:ActiveRole")
	// 角色--用户
	beego.Router("/main/role/to_role_user_add", &auth.RoleController{}, "get:ToRoleUser")
	beego.Router("/main/role/do_role_user_add", &auth.RoleController{}, "post:DoRoleUser")

	// 角色--权限
	beego.Router("/main/role/to_role_auth_add", &auth.RoleController{}, "get:ToRoleAuth")
	beego.Router("/main/role/get_auth_json", &auth.RoleController{}, "get:GetAuthJson")
	beego.Router("/main/role/do_role_auth_add", &auth.RoleController{}, "post:DoRoleAuth")

	// 个人中心
	beego.Router("/main/user/my_center", &user.MyCenterController{})
	beego.Router("/main/user/salary_slip", &user.SalarySlipController{})
	beego.Router("/main/user/salary_slip_detail", &user.SalarySlipController{}, "get:Detail")

	// 财务中心
	beego.Router("/main/caiwu/salary_slip_list", &caiwu.CaiwuEchartDataController{})
	beego.Router("/main/caiwu/to_echart_data_import", &caiwu.CaiwuEchartDataController{}, "get:ToImportExcel")
	beego.Router("/main/caiwu/do_echart_data_import", &caiwu.CaiwuEchartDataController{}, "post:DoImportExcel")
	beego.Router("/main/caiwu/salary_slip_list", &caiwu.CaiwuSalarySlipController{})
	beego.Router("/main/caiwu/to_salary_slip_import", &caiwu.CaiwuSalarySlipController{}, "get:ToImportExcel")
	beego.Router("/main/caiwu/do_salary_slip_import", &caiwu.CaiwuSalarySlipController{}, "post:DoImportExcel")

	// 内容管理
	beego.Router("/main/news/category_list", &news.CategoryController{})
	beego.Router("/main/news/to_add_category", &news.CategoryController{}, "get:ToAdd")
	beego.Router("/main/news/do_add_category", &news.CategoryController{}, "post:DoAdd")
	beego.Router("/main/news/news_list", &news.NewsController{})
	beego.Router("/main/news/to_news_audit", &news.NewsController{}, "get:ToAdd")
	beego.Router("/main/news/do_news_audit", &news.NewsController{}, "post:DoAdd")
	beego.Router("/main/news/upload_img", &news.NewsController{}, "post:UploadImg")
	beego.Router("/main/news/to_edit", &news.NewsController{}, "get:ToEdit")
	beego.Router("/main/news/do_edit", &news.NewsController{}, "post:DoEdit")

	// 车辆管理
	beego.Router("/main/cars/car_brand_list", &cars.CarBrandController{})
	beego.Router("/main/cars/to_car_brand_add", &cars.CarBrandController{}, "get:ToAdd")
	beego.Router("/main/cars/do_car_brand_add", &cars.CarBrandController{}, "post:DoAdd")

	beego.Router("/main/cars/cars_list", &cars.CarsController{})
	beego.Router("/main/cars/to_cars_add", &cars.CarsController{}, "get:ToAdd")
	beego.Router("/main/cars/do_cars_add", &cars.CarsController{}, "post:DoAdd")

	beego.Router("/main/cars/cars_apply_list", &cars.CarsApplyController{})
	beego.Router("/main/cars/to_cars_apply", &cars.CarsApplyController{}, "get:ToApply")
	beego.Router("/main/cars/do_cars_apply", &cars.CarsApplyController{}, "post:DoApply")
	beego.Router("/main/cars/my_apply", &cars.CarsApplyController{}, "get:MyApply")
	beego.Router("/main/cars/audit_apply", &cars.CarsApplyController{}, "get:AuditApply")
	beego.Router("/main/cars/to_audit_apply", &cars.CarsApplyController{}, "get:ToAuditApply")
	beego.Router("/main/cars/do_audit_apply", &cars.CarsApplyController{}, "post:DoAuditApply")
	beego.Router("/main/cars/do_return", &cars.CarsApplyController{}, "get:DoReturn")

}
