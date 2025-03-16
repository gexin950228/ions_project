package login

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_project/models/auth"
	"ions_project/utils"
)

type LController struct {
	beego.Controller
}

func (c *LController) Get() {
	id, base64, err := utils.GetCaptcha()
	if err != nil {
		logs.Error(err)
	}
	c.Data["captcha"] = utils.Captcha{Id: id, BS64: base64}
	c.TplName = "login/login.html"
}

func (c *LController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	captcha := c.GetString("captcha")
	captchaId := c.GetString("captcha_id")

	md5Pass := utils.GetMd5Str(password)

	userInfo := auth.User{}

	o := orm.NewOrm()
	isExist := o.QueryTable("sys_user").Filter("user_name", username).Exist()
	err := o.QueryTable("sys_user").Filter("username", username).Filter("password", md5Pass).One(&userInfo)
	if err != nil {
		logs.Error(err)
	}

	// 验证码校验
	isOk := utils.VerifyCaptcha(captchaId, captcha)
	retMap := map[string]interface{}{}
	if !isExist {
		ret := fmt.Sprintf("登录的post请求，用户名密码错误，登录信息：username:%s;pwd:%s", username, md5Pass)
		logs.Info(ret)
		retMap["code"] = 10001
		retMap["msg"] = "用户名或者密码错误"
		c.Data["json"] = retMap
	} else if !isOk {
		ret := fmt.Sprintf("登录的post请求验证码错误，验证码信息:%v", captcha)
		logs.Info(ret)
		retMap["code"] = 10002
		retMap["msg"] = "验证码错误"
		c.Data["json"] = retMap
	} else if userInfo.IsActive == 0 {
		ret := fmt.Sprintf("登录的post请求，该用户已停用，用户名：%s，状态：停用", username)
		logs.Info(ret)
		retMap["code"] = 10003
		retMap["msg"] = "该用户已停用，请联系管理员"
		c.Data["json"] = retMap
	} else {
		ret := fmt.Sprintf("登录的post请求，登录成功，登录信息：username:%s;pwd:%s", username, md5Pass)
		logs.Info(ret)
		c.SetSession("id", userInfo.Id)
		retMap["code"] = "200"
		retMap["msg"] = "登录成功"
		c.Data["json"] = retMap
	}
	c.ServeJSON()
}

func (c *LController) ChangeCaptcha() {
	message := map[string]string{}
	id, base64, err := utils.GetCaptcha()
	if err != nil {
		ret := fmt.Sprintf("生成验证码失败，错误信息：%s", err.Error())
		logs.Error(ret)
		message["code"] = "10001"
		message["msg"] = err.Error()
		c.Data["json"] = message
	} else {
		c.Data["json"] = utils.Captcha{Id: id, BS64: base64, Code: 200}
	}
	c.ServeJSON()
}

func (c *LController) LogOut() {
	c.DelSession("id")
	c.Redirect(beego.URLFor("LController.Get"), 302)
}
