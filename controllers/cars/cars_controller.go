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

type CarsController struct {
	beego.Controller
}

func (c *CarsController) Get() {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_cars")
	var carsData []auth.Cars

	pagePerNum := 3
	currentPage, err := c.GetInt("page")
	if err != nil { // 说明没有获取到当前页
		currentPage = 1
	}

	offsetNum := (currentPage - 1) * pagePerNum

	var errMsg string
	kw := c.GetString("kw")
	var count int64
	ret := fmt.Sprintf("当前页;%d,查询条件：%s", currentPage, kw)
	logs.Info(ret)
	if kw != "" { // 有查询条件的
		// 总数
		count, _ = qs.Filter("is_delete", 0).Filter("name__contains", kw).Count()
		_, err := qs.Filter("is_delete", 0).Filter("name__contains", kw).Limit(pagePerNum).Offset(offsetNum).RelatedSel().All(&carsData)
		if err != nil {
			errMsg += err.Error()
		}
	} else {
		count, _ = qs.Filter("is_delete", 0).Count()
		_, err := qs.Filter("is_delete", 0).Limit(pagePerNum).Offset(offsetNum).RelatedSel().All(&carsData)
		if err != nil {
			errMsg += err.Error()
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
	if errMsg != "" {
		logs.Error(errMsg)
	}
	pageMap := utils.Paginator(currentPage, pagePerNum, count)
	c.Data["cars_data"] = carsData
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["currentPage"] = currentPage
	c.Data["countPage"] = countPage
	c.Data["count"] = count
	c.Data["page_map"] = pageMap
	c.Data["kw"] = kw
	c.TplName = "cars/cars_list.html"
}

func (c *CarsController) ToAdd() {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_cars_brand")
	var carsBrand []auth.CarBrand
	_, err := qs.Filter("is_delete", 0).All(&carsBrand)
	if err != nil {
		logs.Error(err.Error())
	}
	c.Data["cars_brand"] = carsBrand
	c.TplName = "cars/cars_add.html"
}

func (c *CarsController) DoAdd() {
	o := orm.NewOrm()
	carsBrandId, _ := c.GetInt("cars_brand_id")
	name := c.GetString("name")
	isActive, _ := c.GetInt("is_active")

	carsBrand := auth.CarBrand{Id: carsBrandId}
	carsData := auth.Cars{Name: name, IsActive: isActive, CarBrand: &carsBrand}
	_, err := o.Insert(&carsData)

	messageMap := map[string]interface{}{}
	if err != nil {
		messageMap["code"] = 10001
		messageMap["msg"] = "添加失败"
	}
	messageMap["code"] = 200
	messageMap["msg"] = "添加成功"
	c.Data["json"] = messageMap
	c.ServeJSON()
}
