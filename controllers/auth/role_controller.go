package auth

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_project/models/auth"
	"ions_project/utils"
	"math"
	"strings"
	"time"
)

type RoleController struct {
	beego.Controller
}

func (c *RoleController) List() {
	o := orm.NewOrm()
	var prePage, nextPage, count, currentPage, offsetNum, rolePerPage, countPage int64
	var roles []auth.Role
	rolePerPage = 3
	current, err := c.GetInt("page")
	if err != nil {
		currentPage = 1
		logs.Error(err)
		offsetNum = (currentPage - 1) * rolePerPage
	} else {
		currentPage = int64(current)
		offsetNum = (currentPage - 1) * rolePerPage
	}
	kw := c.GetString("kw")
	ret := fmt.Sprintf("当前页;%d,查询条件：%s", currentPage, kw)
	logs.Info(ret)

	qs := o.QueryTable("sys_role")
	if kw != "" {
		count, _ = qs.Filter("is_delete", 0).Filter("role_name__contains", kw).Count()
		_, err := qs.Filter("is_delete", 0).Filter("role_name__contains", kw).Limit(rolePerPage).Offset(offsetNum).All(&roles)
		if err != nil {
			logs.Error(err)
		}
	} else {
		count, _ = qs.Filter("is_delete", 0).Count()
		_, err = qs.Filter("is_delete", 0).Limit(rolePerPage).Offset(offsetNum).All(&roles)
		ret := fmt.Sprintf("当前页;%d,查询条件：%s", currentPage, kw)
		logs.Info(ret)
	}
	countPage = int64(math.Ceil(float64(count) / float64(rolePerPage)))
	prePage = 1
	if currentPage == 1 {
		prePage = currentPage
	} else {
		prePage = int64(currentPage) - 1
	}
	nextPage = 1
	if currentPage < countPage {
		nextPage = currentPage + 1
	} else {
		nextPage = currentPage
	}
	if err != nil {
		logs.Error(fmt.Sprintf("查询用户列表出错，错误信息: %s", err.Error()))
	}

	countPage = int64(math.Ceil(float64(count) / float64(rolePerPage)))
	prePage = 1
	if currentPage == 1 {
		prePage = currentPage
	} else {
		prePage = currentPage - 1
	}
	nextPage = 1
	if currentPage < countPage {
		nextPage = currentPage + 1
	} else {
		nextPage = currentPage
	}
	page_map := utils.Paginator(int(currentPage), int(rolePerPage), count)

	c.Data["prePage"] = prePage
	c.Data["page_map"] = page_map
	c.Data["nextPage"] = nextPage
	c.Data["count"] = count
	c.Data["lastpage"] = countPage
	c.Data["countPage"] = countPage
	c.Data["roles"] = roles
	c.Data["kw"] = kw
	c.TplName = "auth/role_list.html"
}

func (c *RoleController) ToAdd() {
	c.TplName = "auth/role_add.html"
}

func (c *RoleController) DoAdd() {
	role_name := c.GetString("role_name")
	desc := c.GetString("desc")
	is_active, _ := c.GetInt("is_active")
	role := auth.Role{RoleName: role_name, Desc: desc, CreateTime: time.Now(), IsActive: is_active}
	o := orm.NewOrm()
	_, err := o.Insert(&role)
	message_map := map[string]interface{}{}
	if err != nil {
		message_map["code"] = 10001
		message_map["msg"] = "添加数据错误，请重新添加"
	} else {
		message_map["code"] = 200
		message_map["msg"] = "添加成功"
	}
	c.Data["json"] = message_map
	c.ServeJSON()
}

// ToRoleUser 给用户分配角色
func (c *RoleController) ToRoleUser() {
	uid := c.GetSession("id")
	userId := uid.(int)
	if userId != 1 {
		c.TplName = "auth/403.html"
	}
	id, err := c.GetInt("role_id")
	if err != nil {
		id = 1
	}
	o := orm.NewOrm()
	role := auth.Role{}
	o.QueryTable("sys_role").Filter("id", id).One(&role)

	// 已绑定的用户
	_, err = o.LoadRelated(&role, "User")

	if err != nil {
		logs.Error(err.Error())
	}
	// 未绑定的用户
	var users []auth.User
	if len(role.User) > 0 {
		_, err := o.QueryTable("sys_user").Filter("is_delete", 0).Exclude("id__in", role.User).All(&users)
		if err != nil {
			logs.Error(err)
		}
	}
	c.Data["users"] = users
	c.Data["role"] = role
	c.TplName = "auth/role-user-add.html"
}

func (c *RoleController) DoRoleUser() {
	role_id, _ := c.GetInt("role_id")
	user_ids := c.GetString("user_ids")

	//new_user_ids := user_ids[1:len(user_ids)-1]
	user_id_arr := strings.Split(user_ids, ",")

	// "10,12,13"

	o := orm.NewOrm()
	role := auth.Role{Id: role_id}

	// 查询出已绑定的数据
	m2m := o.QueryM2M(&role, "User")
	m2m.Clear()
	for _, user_id := range user_id_arr {
		user := auth.User{Id: utils.StrToInt(user_id)}
		m2m := o.QueryM2M(&role, "User")
		m2m.Add(user)
	}
	c.Data["json"] = map[string]interface{}{"code": 200, "msg": "添加成功"}
	c.ServeJSON()
}

func (c *RoleController) ActiveRole() {
	role_id, _ := c.GetInt("role_id")
	var errMap []string
	var role auth.Role
	o := orm.NewOrm()
	qs := o.QueryTable("sys_role")
	err := qs.Filter("id", role_id).One(&role)
	if err != nil {
		errMap = append(errMap, err.Error())
		logs.Error(err.Error())
	}
	if role.IsActive == 1 {
		role.IsActive = 0
	} else {
		role.IsActive = 1
	}
	_, err = o.Update(&role)

	if err != nil {
		errMap = append(errMap, err.Error())
		logs.Error(err.Error())
	}
	rep := make(map[string]string)
	rep["code"] = "200"
	rep["msg"] = "操作成功"
	if len(errMap) > 0 {
		rep["code"] = "10001"
		rep["msg"] = "操作失败"
	} else {
		if role.IsActive == 1 {
			rep["code"] = "200"
			rep["msg"] = "启用成功"
		} else {
			rep["code"] = "200"
			rep["msg"] = "停用成功"
		}
	}
	c.Data["json"] = rep
	c.ServeJSON()
}

// 角色，用户配置

func (c *RoleController) ToRoleAuth() {
	role_id, _ := c.GetInt("role_id")
	if role_id == 0 {
		role_id = 1
	}
	role := auth.Role{}
	o := orm.NewOrm()
	qs := o.QueryTable("sys_role")
	qs.Filter("id", role_id).RelatedSel().One(&role)
	c.Data["role"] = role
	c.TplName = "auth/role-auth-add.html"
}

func (c *RoleController) GetAuthJson() {
	roleId, _ := c.GetInt("role_id")

	o := orm.NewOrm()
	qs := o.QueryTable("sys_auth")
	role := auth.Role{Id: roleId}
	_, err := o.LoadRelated(&role, "Auth")
	if err != nil {
		logs.Error(err)
	}

	var authIdsHas []int
	for _, authData := range role.Auth {
		authIdsHas = append(authIdsHas, authData.Id)
	}

	//	 所有权限
	var auths []auth.Auth
	_, err = qs.Filter("is_delete", 0).All(&auths)
	if err != nil {
		logs.Error(err.Error())
	}

	var authArrMap []map[string]interface{}
	for _, authData := range auths {
		id := authData.Id
		pId := authData.Pid
		name := authData.AuthName
		if pId == 0 {
			authMap := map[string]interface{}{"id": id, "pId": pId, "name": name, "open": false}
			authArrMap = append(authArrMap, authMap)
		} else {
			authMap := map[string]interface{}{"id": id, "pId": pId, "name": name}
			authArrMap = append(authArrMap, authMap)
		}
	}
	authMaps := map[string]interface{}{}
	authMaps["auth_arr_map"] = authArrMap
	authMaps["auth_ids_has"] = authIdsHas
	c.Data["json"] = authMaps
	c.ServeJSON()
}

func (c *RoleController) DoRoleAuth() {
	role_id, _ := c.GetInt("role_id")
	auth_ids := c.GetString("auth_ids")
	id_arr := strings.Split(auth_ids, ",")

	o := orm.NewOrm()
	role := auth.Role{Id: role_id}
	m2m := o.QueryM2M(&role, "Auth")
	m2m.Clear()

	for _, auth_id := range id_arr {
		auth_id_int := utils.StrToInt(auth_id)
		if auth_id_int != 0 {
			auth_data := auth.Auth{Id: auth_id_int}
			m2m := o.QueryM2M(&role, "Auth")
			m2m.Add(&auth_data)
		}
	}
	c.Data["json"] = map[string]interface{}{"code": 200, "msg": "添加成功"}
	c.ServeJSON()
}
