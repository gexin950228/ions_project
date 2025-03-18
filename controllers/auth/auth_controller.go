package auth

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ions_project/models/auth"
	"ions_project/utils"
	"math"
)

type AuthController struct {
	beego.Controller
}

func (a *AuthController) List() {
	pagePerNum := 3
	currentPage, err := a.GetInt("page")
	if err != nil {
		currentPage = 1
	}
	o := orm.NewOrm()
	qs := o.QueryTable("sys_auth")
	count, _ := qs.Filter("is_delete", 0).Count()
	countPage := int(math.Ceil(float64(count) / float64(pagePerNum)))
	nextPage := 1
	if currentPage < countPage {
		nextPage = currentPage + 1
	} else if currentPage >= countPage {
		nextPage = countPage
	}
	// 关键字查询
	kw := a.GetString("kw")
	if kw == "" {

	}
	offset := (currentPage - 1) * pagePerNum
	page_map := utils.Paginator(currentPage, pagePerNum, count)
	var auths []auth.Auth
	qs1 := o.QueryTable("sys_auth")
	qs1.Filter("is_delete", 0).Offset(offset).Limit(pagePerNum).All(&auths)
	for _, sAuth := range auths {
		pid := sAuth.Pid
		var pAuth auth.Auth
		qs2 := o.QueryTable("sys_auth")
		qs2.Filter("id", pid).One(&pAuth)
		sAuth.PName = pAuth.AuthName
	}
	a.Data["auths"] = auths
	a.Data["nextPage"] = nextPage
	a.Data["page_map"] = page_map
	a.Data["countPage"] = countPage
	a.Data["count"] = count
	a.TplName = "auth/auth-list.html"
}

func (a *AuthController) ToAuthAdd() {
	var auths []auth.Auth
	o := orm.NewOrm()
	qs := o.QueryTable("sys_auth")
	qs.Filter("is_delete", 0).All(&auths)
	a.Data["auths"] = auths
	a.TplName = "auth/role_list.html"
}

func (a *AuthController) DoAdd() {
	var repInfo map[string]interface{}
	pid := 1
	var errArr []error
	authParentId, err := a.GetInt("auth_parent_id")
	if err != nil {
		pid = 0
	} else {
		pid = authParentId
	}
	authName := a.GetString("auth_name")
	authUrl := a.GetString("auth_url")
	authDesc := a.GetString("auth_desc")
	isActive, err := a.GetInt("is_active")
	if err != nil {
		errArr = append(errArr, err)
		isActive = 1
	}
	authWeight, err := a.GetInt("auth_weight")
	if err != nil {
		authWeight = 1
		errArr = append(errArr, err)
	}
	var auth auth.Auth
	auth.AuthName = authName
	auth.Pid = pid
	auth.Desc = authDesc
	auth.UrlFor = authUrl
	auth.Weight = authWeight
	auth.IsActive = isActive
	o := orm.NewOrm()
	_, err = o.Insert(&auth)
	if err != nil {
		errArr = append(errArr, err)
	}
	if len(errArr) > 0 {
		repInfo["code"] = "500"
		repInfo["msg"] = errArr
	} else {
		repInfo["code"] = "200"
		repInfo["msg"] = "添加权限成功"
	}
	a.Data["json"] = repInfo
	a.ServeJSON()
}
