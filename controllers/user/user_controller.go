package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_project/models/auth"
	"ions_project/utils"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) List() {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_user")
	var users []auth.User
	pagePerNum := 5
	currentPage, err := u.GetInt("page")
	if err != nil {
		logs.Error(err)
		currentPage = 1
	}
	offsetNum := (currentPage - 1) * pagePerNum
	kw := u.GetString("kw")
	var count int64 = 0
	ret := fmt.Sprintf("当前页;%d,查询条件：%s", currentPage, kw)
	logs.Info(ret)
	if kw != "" {
		count, _ = qs.Filter("is_delete", 0).Filter("user_name__contains", kw).Count()
		_, err := qs.Filter("is_delete", 0).Filter("user_name__contains", kw).Limit(pagePerNum).Offset(offsetNum).All(&users)
		if err != nil {
			logs.Error(err)
		}
	} else {
		count, _ = qs.Filter("is_delete", 0).Count()
		_, err := qs.Filter("is_delete", 0).Limit(pagePerNum).Offset(offsetNum).All(&users)
		if err != nil {
			logs.Error(fmt.Sprintf("查询用户列表出错，错误信息: %s", err.Error()))
		}
	}
	countPage := int(math.Ceil(float64(count) / float64(pagePerNum)))
	prePage := 1
	if currentPage == 1 {
		prePage = currentPage
	} else {
		prePage = currentPage - 1
	}
	nextPage := 1
	if currentPage < countPage {
		nextPage = currentPage + 1
	} else {
		nextPage = currentPage
	}
	pageMap := utils.Paginator(currentPage, pagePerNum, count)
	u.Data["users"] = users
	u.Data["kw"] = kw
	u.Data["count"] = count
	u.Data["page_map"] = pageMap
	u.Data["nextPage"] = nextPage
	u.Data["currentPage"] = currentPage
	u.Data["prePage"] = prePage
	u.Data["countPage"] = countPage
	u.TplName = "user/user_list.html"
}

func (u *UserController) ToAdd() {
	u.TplName = "user/user_add.html"
}

func (u *UserController) DoAdd() {
	fmt.Println("进入")
	username := u.GetString("username")
	password := u.GetString("password")
	age, _ := u.GetInt("age")
	gender, _ := u.GetInt("gender")
	phone := u.GetString("phone")
	addr := u.GetString("addr")
	is_active, _ := u.GetInt("is_active")
	new_password := utils.GetMd5Str(password)
	phoneInt64, _ := strconv.ParseInt(phone, 10, 64)
	o := orm.NewOrm()
	userData := auth.User{UserName: username, Password: new_password, Age: age, Gender: gender, Phone: phoneInt64, Addr: addr, IsActive: is_active}
	_, err := o.Insert(&userData)
	message_map := make(map[string]interface{})
	fmt.Println(age, username, gender, addr, new_password, phoneInt64, is_active)
	if err != nil {
		ret1 := fmt.Sprintf("插入数据信息：username:%s|md5_password:%s|age:%d|gender:%d|phone:%s|"+
			"addr:%s;is_active:%d", username, new_password, age, gender, phone, addr, is_active)
		ret := fmt.Sprintf("添加数据出错,错误信息:%v", err)
		logs.Error(ret1)
		logs.Error(ret)
		message_map["code"] = 10001
		message_map["msg"] = "添加数据出错，请重新添加"
		u.Data["json"] = message_map
	} else {
		ret1 := fmt.Sprintf("插入数据成功，数据信息：username:%s|md5_password:%s|age:%d|gender:%d|phone:%s|"+
			"addr:%s;is_active:%d", username, new_password, age, gender, phone, addr, is_active)
		logs.Info(ret1)
		message_map["code"] = 200
		message_map["msg"] = "添加成功"
		u.Data["json"] = message_map
	}
	u.ServeJSON()
}

func (u *UserController) IsActive() {
	is_active, _ := u.GetInt("is_active_val")
	id, _ := u.GetInt("id")

	o := orm.NewOrm()
	qs := o.QueryTable("sys_user").Filter("id", id)
	message_map := map[string]interface{}{}
	if is_active == 1 {
		_, err := qs.Update(orm.Params{"is_active": 0})
		if err != nil {
			logs.Error(fmt.Sprintf("更新用户信息出错，错误信息:%s", err.Error()))
		}
		ret := fmt.Sprintf("用户id:%d,停用成功", id)
		logs.Info(ret)
		message_map["msg"] = "停用成功"
	} else if is_active == 0 {
		_, err := qs.Update(orm.Params{"is_active": 1})
		if err != nil {
			logs.Error(fmt.Sprintf("更新用户信息出错，错误信息:%s", err.Error()))
		}
		ret := fmt.Sprintf("用户id:%d,启用成功", id)
		logs.Info(ret)
		message_map["msg"] = "启用成功"
	}
	u.Data["json"] = message_map
	u.ServeJSON()
}

func (u *UserController) Delete() {
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	_, err := o.QueryTable("sys_user").Filter("id", id).Update(orm.Params{"is_delete": 1})
	if err != nil {
		logs.Error("id为%s的用户删除出错，错误信息为:%v", id, err.Error())
	}
	ret := fmt.Sprintf("用户id:%d,删除成功", id)
	logs.Info(ret)
	u.Redirect(beego.URLFor("UserController.List"), http.StatusFound)
}

func (u *UserController) ResetPassword() {
	id, _ := u.GetInt("id")

	o := orm.NewOrm()
	newPassword := utils.GetMd5Str("Admin@123456")
	_, err := o.QueryTable("sys_user").Filter("id", id).Update(orm.Params{"password": newPassword})
	if err != nil {
		logs.Error(fmt.Sprintf("更新id为%d的用户密码发生错误，错误信息为: %s", id, err.Error()))
	}
	ret := fmt.Sprintf("用户id:%d,重置密码成功", id)
	logs.Info(ret)
	u.Redirect(beego.URLFor("UserController.List"), http.StatusFound)
}

func (u *UserController) ToUpdate() {
	userId := u.GetSession("id")
	var uid int
	if userId == nil {
		uid = 1
	} else {
		fmt.Println(userId)
		uid = userId.(int)
	}
	o := orm.NewOrm()
	userData := auth.User{}
	err := o.QueryTable("sys_user").Filter("id", uid).One(&userData)
	fmt.Printf("user_data:%v\n", userData)
	if err != nil {
		logs.Error(fmt.Sprintf("查询id为%d的用户出错，错误信息: %s", uid, err.Error()))
	}
	u.Data["user"] = userData
	u.TplName = "user/user_edit.html"
}

func (u *UserController) DoUpdate() {
	uid, _ := u.GetInt("uid")
	username := u.GetString("username")
	password := u.GetString("password")
	age, _ := u.GetInt("age")
	gender, _ := u.GetInt("gender")
	phone := u.GetString("phone")
	addr := u.GetString("addr")
	is_active, _ := u.GetInt("is_active")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_user").Filter("id", uid)
	message_map := map[string]interface{}{}
	new_password := utils.GetMd5Str(password)
	if password == "" {
		_, err := qs.Update(orm.Params{
			"username":  username,
			"age":       age,
			"gender":    gender,
			"phone":     phone,
			"addr":      addr,
			"is_active": is_active,
		})
		if err != nil {
			ret := fmt.Sprintf("更新失败，用户id:%d", uid)
			logs.Error(ret)
			message_map["code"] = 10001
			message_map["msg"] = "更新失败"
		} else {
			ret := fmt.Sprintf("更新成功，用户id:%d", uid)
			logs.Info(ret)
			message_map["code"] = 200
			message_map["msg"] = "更新成功"
		}
	} else {
		_, err := qs.Update(orm.Params{
			"username":  username,
			"password":  new_password,
			"age":       age,
			"gender":    gender,
			"phone":     phone,
			"addr":      addr,
			"is_active": is_active,
		})

		if err != nil {

			ret := fmt.Sprintf("更新失败，用户id:%d", uid)
			logs.Error(ret)
			message_map["code"] = 10001
			message_map["msg"] = "更新失败"
		} else {
			ret := fmt.Sprintf("更新成功，用户id:%d", uid)
			logs.Info(ret)
			message_map["code"] = 200
			message_map["msg"] = "更新成功"
		}
	}

	u.Data["json"] = message_map
	u.ServeJSON()
}

func (u *UserController) MuliDelete() {
	ids := u.GetString("ids")
	newIds := ids[1 : len(ids)-1]
	idArr := strings.Split(newIds, ",")

	var user auth.User
	o := orm.NewOrm()
	qs := o.QueryTable("sys_user")
	var errArr []string
	for _, id := range idArr {
		idInt, _ := strconv.Atoi(id)
		err := qs.Filter("id", idInt).One(&user)
		if err != nil {
			errArr = append(errArr, err.Error())
			logs.Error(fmt.Sprintf("查询id为%d的用户失败，错误信息: %s", idInt, err.Error()))
		}
		_, err = qs.Filter("id", idInt).Update(orm.Params{"is_delete": 1})
		if err != nil {
			logs.Error(fmt.Sprintf("删除id为%d的用户出错，错误信息: %s", idInt, err.Error()))
			errArr = append(errArr, err.Error())
		}
	}
	var repInfo map[string]interface{}
	if len(errArr) > 0 {
		repInfo["code"] = "500"
		repInfo["msg"] = errArr
	} else {
		repInfo["code"] = "200"
		repInfo["msg"] = "删除成功"
	}
	u.Data["json"] = repInfo
	u.ServeJSON()
}
