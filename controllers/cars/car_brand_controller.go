package cars

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_project/models/auth"
	"ions_project/utils"
	"math"
)

type CarBrandController struct {
	beego.Controller
}

func (c *CarBrandController) Get() {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_cars_brand")
	var carsBrand []auth.CarBrand
	pagePerNum := 3
	currentPage, err := c.GetInt("page")
	if err != nil {
		currentPage = 1
	}
	var errorInfo string
	offsetNum := (currentPage - 1) * pagePerNum
	kw := c.GetString("kw")
	var count int64
	ret := fmt.Sprintf("当前页码:%d，查询条件: %s", currentPage, kw)
	logs.Info(ret)
	if kw == "" {
		count, _ = qs.Filter("is_delete", 0).Count()
		_, err := qs.Filter("is_delete", 0).Limit(pagePerNum).Offset(offsetNum).All(&carsBrand)
		if err != nil {
			logs.Error(fmt.Sprintf("查询车牌信息出错，错误信息: %s", err.Error()))
			errorInfo += err.Error()
		}
	} else {
		count, _ = qs.Filter("is_delete", 0).Filter("name__contains", kw).Count()
		_, err := qs.Filter("is_delete", 0).Filter("name__contains", kw).All(&carsBrand)
		if err != nil {
			logs.Error(err)
			errorInfo += err.Error()
		}
	}

	countPage := int(math.Ceil(float64(count) / float64(pagePerNum)))

	var prePage, nextPage int
	if currentPage == 1 {
		prePage = currentPage
	} else {
		prePage = currentPage - 1
	}

	if currentPage < countPage {
		nextPage = currentPage + 1
	} else {
		nextPage = currentPage
	}
	page_map := utils.Paginator(currentPage, pagePerNum, count)
	c.Data["cars_brand"] = carsBrand
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["currentPage"] = currentPage
	c.Data["countPage"] = countPage
	c.Data["count"] = count
	c.Data["page_map"] = page_map
	c.Data["kw"] = kw
	c.TplName = "cars/cars_brand_list.html"
}

func (c *CarBrandController) ToAdd() {
	c.TplName = "cars/cars_brand_add.html"
}

func (c *CarBrandController) DoAdd() {
	name := c.GetString("name")
	desc := c.GetString("desc")
	isActive, _ := c.GetInt("is_active")
	o := orm.NewOrm()
	carsBrand := auth.CarBrand{
		Name:     name,
		Desc:     desc,
		IsActive: isActive,
	}
	var msgMap map[string]interface{}
	_, err := o.Insert(&carsBrand)
	if err != nil {
		msgMap["code"] = 10001
		msgMap["msg"] = "添加失败"
	}
	msgMap["code"] = 200
	msgMap["msg"] = "添加成功"

	c.Data["json"] = msgMap
	c.ServeJSON()
}
