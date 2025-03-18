package news

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_project/models/news"
	"ions_project/utils"
	"math"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_category")
	var categories []news.Category
	pagePerNum := 5
	currentPage, err := c.GetInt("page")
	if err != nil {
		currentPage = 1
	}

	offsetNum := (currentPage - 1) * pagePerNum
	kw := c.GetString("kw")
	var count int64 = 0
	ret := fmt.Sprintf("当前页;%d,查询条件：%s", currentPage, kw)
	logs.Info(ret)
	if kw != "" { // 有查询条件的
		// 总数
		count, _ = qs.Filter("is_delete", 0).Filter("title__contains", kw).Count()
		_, err := qs.Filter("is_delete", 0).Filter("title__contains", kw).Limit(pagePerNum).Offset(offsetNum).RelatedSel().All(&categories)
		if err != nil {
			logs.Error(fmt.Sprintf("查询categoty出错，错误信息: %s", err.Error()))
		}
	} else {
		count, _ = qs.Filter("is_delete", 0).Count()
		_, err := qs.Filter("is_delete", 0).Limit(pagePerNum).Offset(offsetNum).RelatedSel().All(&categories)
		if err != nil {
			logs.Error(fmt.Sprintf("查询categoty出错，错误信息: %s", err.Error()))
		}
	}
	fmt.Println("==================================================")
	fmt.Printf("categories: %v\n", categories)
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
	page_map := utils.Paginator(currentPage, pagePerNum, count)
	c.Data["categories"] = categories
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["currentPage"] = currentPage
	c.Data["countPage"] = countPage
	c.Data["count"] = count
	c.Data["page_map"] = page_map
	c.Data["kw"] = kw

	c.TplName = "news/category_list.html"
}

func (c *CategoryController) ToAdd() {
	c.TplName = "news/category_add.html"
}

func (c *CategoryController) DoAdd() {
	name := c.GetString("name")
	desc := c.GetString("desc")
	isActive, _ := c.GetInt("is_active")
	o := orm.NewOrm()
	category := news.Category{
		Name:     name,
		Desc:     desc,
		IsActive: isActive,
	}
	_, err := o.Insert(&category)
	messageMap := map[string]interface{}{}
	if err != nil {
		messageMap["code"] = 10001
		messageMap["msg"] = "添加栏目失败"

	}

	messageMap["code"] = 200
	messageMap["msg"] = "添加成功"

	c.Data["json"] = messageMap
	c.ServeJSON()
}
