package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_project/models/auth"
	"ions_project/utils"
	"math"
	"time"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	// 后端首页
	orm.Debug = true
	o := orm.NewOrm()
	var uid int
	userId := c.GetSession("id")
	if userId == nil {
		uid = 1
	} else {
		uid = userId.(int)
	}
	user := auth.User{Id: uid}
	o.LoadRelated(&user, "Role")
	var auth_arr []int
	for _, role := range user.Role {
		fmt.Println(role)
		role_data := auth.Role{Id: role.Id}
		o.LoadRelated(&role_data, "Auth")
		fmt.Printf("role_data: %v\n", role_data)
		for _, auth_date := range role_data.Auth {
			auth_arr = append(auth_arr, auth_date.Id)
		}
	}
	qs := o.QueryTable("sys_auth")
	auths := []auth.Auth{}
	qs.Filter("pid", 0).Filter("id__in", auth_arr).OrderBy("-weight").All(&auths)
	//"select * from sys_user where id in (1,2,3,1)"
	trees := []auth.Tree{}
	for _, auth_data := range auths { // 一级菜单
		pid := auth_data.Id // 根据pid获取所有的子解点
		tree_data := auth.Tree{Id: auth_data.Id, AuthName: auth_data.AuthName, UrlFor: auth_data.UrlFor, Weight: auth_data.Weight, Children: []*auth.Tree{}}
		GetChildNode(pid, &tree_data)
		trees = append(trees, tree_data)
	}
	o.QueryTable("sys_user").Filter("id", uid).One(&user)
	qs1 := o.QueryTable("sys_cars_apply")
	cars_apply := []auth.CarsApply{}
	qs1.Filter("user_id", uid).Filter("return_status", 0).Filter("notify_tag", 0).All(&cars_apply)

	cur_time, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	for _, apply := range cars_apply {
		return_date := apply.ReturnDate
		ret := cur_time.Sub(return_date)
		content := fmt.Sprintf("%s用户，你借的车辆归还时间为%v,已经预期，请尽快归还!!", user.UserName, return_date.Format("2006-01-02"))
		if ret > 0 { // 已经逾期
			message_notify := auth.MessageNotify{
				Flag:    1,
				Title:   "车辆归还逾期",
				Content: content,
				User:    &user,
				ReadTag: 0,
			}
			o.Insert(&message_notify)
		}
		apply.NotifyTag = 1
		o.Update(&apply)
	}
	// 展示消息,使用websocket优化
	qs2 := o.QueryTable("sys_message_notify")
	notify_count, _ := qs2.Filter("read_tag", 0).Count()
	fmt.Printf("trees: %v\n", trees)
	c.Data["notify_count"] = notify_count
	c.Data["trees"] = trees
	c.Data["user"] = user
	c.TplName = "index.html"
}

func (c *HomeController) Add() {}

func (c *HomeController) Welcome() {
	c.TplName = "welcome.html"
}

func GetChildNode(pid int, treeNode *auth.Tree) {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_auth")
	var auths []auth.Auth
	_, err := qs.Filter("pid", pid).OrderBy("-weight").All(&auths)
	if err != nil {
		logs.Error(err)
	}

	for i := 0; i < len(auths); i++ {
		pid := auths[i].Id
		treeData := auth.Tree{Id: auths[i].Id, AuthName: auths[i].AuthName, UrlFor: auths[i].UrlFor, Weight: auths[i].Weight, Children: []*auth.Tree{}}
		treeNode.Children = append(treeNode.Children, &treeData)
		GetChildNode(pid, &treeData)
	}
	return

}

func (c *HomeController) NotifyList() {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_message_notify")
	var notifies []auth.MessageNotify
	pagePerNum := 8
	offsetNum := (pagePerNum - 1) * pagePerNum
	currentPage, err := c.GetInt("page")
	if err != nil {
		logs.Error(err)
		currentPage = 1
	}
	kw := c.GetString("kw")
	var count int64
	ret := fmt.Sprintf("当前页;%d,查询条件：%s", currentPage, kw)
	logs.Info(ret)
	if kw != "" {
		count, _ = qs.Filter("title__contains", kw).Count()
		_, err := qs.Filter("title__contains", kw).Limit(pagePerNum).Offset(offsetNum).All(&notifies)
		if err != nil {
			logs.Error(err)
		}
	} else {
		count, _ = qs.Count()
		_, err := qs.Limit(pagePerNum).Offset(offsetNum).All(&notifies)
		if err != nil {
			logs.Error(err)
		}
	}
	countPage := int(math.Ceil(float64(count) / float64(pagePerNum)))
	prePage := 1
	if currentPage == 1 {
		prePage = currentPage
	} else if currentPage > 1 {
		prePage = currentPage - 1
	}
	nextPage := 1
	if currentPage < countPage {
		nextPage = currentPage + 1
	} else if currentPage >= countPage {
		nextPage = currentPage
	}
	pageMap := utils.Paginator(currentPage, pagePerNum, count)
	c.Data["page_map"] = pageMap
	c.Data["count"] = count
	c.Data["prePage"] = prePage
	c.Data["currentPage"] = currentPage
	c.Data["kw"] = kw
	c.Data["nextPage"] = nextPage
	c.TplName = "notify_list.html"
}

func (c *HomeController) ReadNotify() {
	id, _ := c.GetInt("id")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_message_notify")
	_, err := qs.Filter("id", id).Update(orm.Params{"read_tag": 1})
	if err != nil {
		logs.Error(err)
	}
	c.Redirect(beego.URLFor("HomeController.NotifyList"), 302)
}
