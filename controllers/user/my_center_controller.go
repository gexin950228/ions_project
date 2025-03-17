package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_project/models/auth"
)

type MyCenterController struct {
	beego.Controller
}

func (m *MyCenterController) Get() {
	id := m.GetSession("id")

	o := orm.NewOrm()
	qs := o.QueryTable("sys_user")
	var user auth.User
	err := qs.Filter("id", id).One(&user)
	if err != nil {
		logs.Error(fmt.Sprintf("get user err: %s", err.Error()))
	}
	m.Data["user"] = user
	m.TplName = "user/my_center_edit.html"
}

func (m *MyCenterController) Post() {

}
