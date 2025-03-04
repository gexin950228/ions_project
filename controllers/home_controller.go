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
	o := orm.NewOrm()
	userId := c.GetSession("id").(int)
	user := auth.User{Id: userId}

	_, err := o.LoadRelated(&user, "Role")
	if err != nil {
		logs.Error(err)
	}

	var authArr []int
	for _, role := range user.Role {
		roleData := auth.Role{Id: role.Id}
		_, err := o.LoadRelated(&roleData, "Auth")
		if err != nil {
			logs.Error(err)
		}
		for _, authData := range roleData.Auth {
			authArr = append(authArr, authData.Id)
		}
	}
	qs := o.QueryTable("sys_auth")
	var auths []auth.Auth
	_, err = qs.Filter("id__in", authArr).OrderBy("-weight").All(&auths)
	if err != nil {
		logs.Error(err)
	}
	var trees []auth.Tree
	var treeData auth.Tree
	var pid int
	for _, authData := range auths {
		pid = authData.Id
		treeData = auth.Tree{Id: authData.Id, AuthName: authData.AuthName, UrlFor: authData.UrlFor, Weight: authData.Weight, Children: []*auth.Tree{}}
		GetChildNode(pid, &treeData)
		trees = append(trees, treeData)
	}
	err = o.QueryTable("sys_auth").Filter("id", userId).One(&user)
	if err != nil {
		logs.Error(fmt.Sprintf("查询用户信息出错: %s", err.Error()))
	}

	// 消息通知，发送消息，使用定时任务
	qs1 := o.QueryTable("sys_cars_apply")
	var carsApply []auth.CarsApply
	_, err = qs1.Filter("user_id", userId).Filter("return_status", 0).Filter("notify_tag", 0).All(&carsApply)
	curTime, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	for _, apply := range carsApply {
		returnDate := apply.ReturnDate
		ret := curTime.Sub(returnDate)
		content := fmt.Sprintf("%s用户的借用车辆归还时间为:%v，请尽快归还", user.UserName, returnDate.Format("2006-01-02"))
		if ret > 0 { // 已逾期
			messageNotify := auth.MessageNotify{
				Content: content,
				Flag:    1,
				User:    &user,
				ReadTag: 0,
			}
			_, err := o.Insert(&messageNotify)
			if err != nil {
				logs.Error(err)
			}
		}
		apply.NotifyTag = 1
		_, err = o.Update(&apply)
		if err != nil {
			logs.Error(err)
		}
	}
	// 展示消息,使用websocket优化
	qs2 := o.QueryTable("sys_message_notify")
	notifyCount, _ := qs2.Filter("read_tag", 0).Count()
	c.Data["user"] = user
	c.Data["trees"] = trees
	c.Data["notify_count"] = notifyCount
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
