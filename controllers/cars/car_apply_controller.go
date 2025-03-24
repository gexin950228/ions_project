package cars

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

type CarsApplyController struct {
	beego.Controller
}

func (c *CarsApplyController) Get() {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_cars")
	var carsData []auth.Cars
	pagePerNum := 3
	currentPage, err := c.GetInt("page")
	if err != nil {
		currentPage = 1
	}
	offsetNum := (currentPage - 1) * pagePerNum
	kw := c.GetString("kw")
	var count int64
	var errInfo string
	ret := fmt.Sprintf("当前页;%d,查询条件：%s", currentPage, kw)
	logs.Info(ret)
	if kw != "" {
		count, _ = qs.Filter("is_delete", 0).Filter("name__contains", kw).Count()
		_, err := qs.Filter("is_delete", 0).Filter("name__contains", kw).Limit(pagePerNum).Offset(offsetNum).RelatedSel().All(&carsData)
		if err != nil {
			errInfo += err.Error()
		}
	} else {
		count, _ = qs.Filter("is_delete", 0).Count()
		_, err := qs.Filter("is_delete", 0).Limit(pagePerNum).Offset(offsetNum).RelatedSel().All(&carsData)
		if err != nil {
			errInfo += err.Error()
		}
	}
	countPage := int(math.Ceil(float64(count) / float64(pagePerNum)))
	var prePage int
	if currentPage == 1 {
		prePage = 1
	} else {
		prePage = currentPage - 1
	}
	nextPage := 1
	if currentPage < countPage {
		nextPage = currentPage + 1
	} else if currentPage >= countPage {
		nextPage = currentPage
	}
	page_map := utils.Paginator(currentPage, pagePerNum, count)
	c.Data["cars_data"] = carsData
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["currentPage"] = currentPage
	c.Data["countPage"] = countPage
	c.Data["count"] = count
	c.Data["page_map"] = page_map
	c.Data["kw"] = kw

	c.TplName = "cars/cars_apply_list.html"
}

func (c *CarsApplyController) ToApply() {
	id, _ := c.GetInt("id")
	c.Data["id"] = id
	c.TplName = "cars/cars_apply.html"

}

func (c *CarsApplyController) DoApply() {
	reason := c.GetString("reason")
	destination := c.GetString("destination")
	returnDate := c.GetString("return_date")
	returnDateNew, _ := time.Parse("2006-01-02", returnDate)
	carsId, _ := c.GetInt("cars_id")
	uid := c.GetSession("id")
	user := auth.User{Id: uid.(int)}
	carsDate := auth.Cars{Id: carsId}
	o := orm.NewOrm()
	carsApply := auth.CarsApply{
		User:         &user,
		Cars:         &carsDate,
		Reason:       reason,
		Destination:  destination,
		ReturnDate:   returnDateNew,
		ReturnStatus: 0,
		AuditStatus:  3,
		IsActive:     1,
	}
	var errInfo string
	_, err := o.Insert(&carsApply)
	if err != nil {
		errInfo += err.Error()
	}

	_, err = o.QueryTable("sys_cars").Filter("id", carsId).Update(orm.Params{
		"status": 2,
	})
	if err != nil {
		logs.Error(err)
		errInfo = errInfo + err.Error()
	}
	messageMap := map[string]interface{}{}
	if errInfo != "" {
		messageMap["code"] = 10001
		messageMap["msg"] = fmt.Sprintf("申请失败，错误信息%s", errInfo)
	}
	messageMap["code"] = 200
	messageMap["msg"] = "申请成功"
	c.Data["json"] = messageMap
	c.ServeJSON()
}

func (c *CarsApplyController) MyApply() {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_cars_apply")
	carsApply := auth.CarsApply{}
	pagePerNum := 3
	currentPage, err := c.GetInt("page")
	if err != nil {
		currentPage = 1
	}
	offsetNum := (currentPage - 1) * pagePerNum
	kw := c.GetString("kw")
	var count int64
	uid := c.GetSession("id")
	ret := fmt.Sprintf("当前页;%d,查询条件：%s", currentPage, kw)
	logs.Info(ret)
	if kw != "" {
		count, _ = qs.Filter("is_delete", 0).Filter("Cars__name__contains", kw).Filter("user_id", uid.(int)).Count()
		_, err := qs.Filter("is_delete", 0).Filter("Cars__name__contains", kw).Limit(pagePerNum).Offset(offsetNum).RelatedSel().Filter("user_id", uid.(int)).All(&carsApply)
		if err != nil {
			logs.Error(err.Error())
		}
	} else {
		count, _ = qs.Filter("is_delete", 0).Filter("user_id", uid.(int)).Count()
		_, err := qs.Filter("is_delete", 0).Limit(pagePerNum).Filter("user_id", uid.(int)).Offset(offsetNum).RelatedSel().All(&carsApply)
		if err != nil {
			logs.Error(err.Error())
		}
	}
	countPage := int(math.Ceil(float64(count) / float64(pagePerNum)))
	var prePage, nextPage int
	if currentPage == 1 {
		prePage = 1
	} else if currentPage > 1 {
		prePage = currentPage - 1
	}
	if currentPage < countPage {
		nextPage = currentPage + 1
	} else if currentPage >= countPage {
		nextPage = currentPage
	}
	pageMap := utils.Paginator(currentPage, pagePerNum, count)
	c.Data["cars_apply"] = carsApply
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["currentPage"] = currentPage
	c.Data["countPage"] = countPage
	c.Data["count"] = count
	c.Data["page_map"] = pageMap
	c.Data["kw"] = kw
	c.TplName = "cars/my_apply_list.html"
}

func (c *CarsApplyController) AuditApply() {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_cars_apply")
	var carsApply []auth.CarsApply
	pagePerNum := 3
	currentPage, err := c.GetInt("page")
	if err != nil { // 说明没有获取到当前页
		currentPage = 1
	}
	offsetNum := pagePerNum * (currentPage - 1)
	kw := c.GetString("kw")
	var count int64
	ret := fmt.Sprintf("当前页;%d,查询条件：%s", currentPage, kw)
	logs.Info(ret)
	var errorInfo []string
	fmt.Println("=================================================")
	fmt.Printf("kw: %s", kw)
	if kw != "" {
		count, _ = qs.Filter("is_delete", 0).Filter("cars_name__contains", kw).Count()
		_, err := qs.Filter("is_delete", 0).Filter("cars_name__contains", kw).Limit(pagePerNum).Offset(offsetNum).All(&carsApply)
		if err != nil {
			errorInfo = append(errorInfo, err.Error())
		}
	} else {
		count, _ = qs.Filter("is_delete", 0).Count()
		_, err := qs.Filter("is_delete", 0).Limit(pagePerNum).All(&carsApply)
		if err != nil {
			errorInfo = append(errorInfo, err.Error())
		}
	}
	fmt.Println("=================================================")
	for _, v := range carsApply {
		fmt.Printf("车辆租借信息: %v\n", v)
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
	fmt.Println("=================================================")
	fmt.Println(carsApply)
	page_map := utils.Paginator(currentPage, pagePerNum, count)
	c.Data["cars_apply"] = carsApply
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["currentPage"] = currentPage
	c.Data["countPage"] = countPage
	c.Data["count"] = count
	c.Data["page_map"] = page_map
	c.Data["kw"] = kw
	c.TplName = "cars/audit_apply_list.html"
}

func (c *CarsApplyController) ToAuditApply() {
	id, _ := c.GetInt("id")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_cars_apply")
	carsApply := auth.CarsApply{}
	err := qs.Filter("id", id).One(&carsApply)
	if err != nil {
		logs.Error(err.Error())
	}
	fmt.Println("====================================================")
	fmt.Printf("carsApply:%#v", carsApply)
	fmt.Printf("id: %v", id)
	c.Data["cars_apply"] = carsApply
	c.TplName = "cars/audit_apply.html"
}

func (c *CarsApplyController) DoAuditApply() {
	option := c.GetString("option")
	auditStatus, _ := c.GetInt("audit_status")
	id, _ := c.GetInt("id")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_cars_apply")
	_, err := qs.Filter("id", id).Update(orm.Params{
		"audit_status": auditStatus,
		"audit_option": option,
	})
	messageMap := map[string]interface{}{}
	if err != nil {
		messageMap["code"] = 10001
		messageMap["msg"] = "审核失败"
	}

	messageMap["code"] = 200
	messageMap["msg"] = "审核成功"

	c.Data["json"] = messageMap
	c.ServeJSON()
}

func (c *CarsApplyController) DoReturn() {
	id, _ := c.GetInt("id")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_cars_apply")
	_, err := qs.Filter("id", id).Update(orm.Params{
		"return_status": 1,
	})
	if err != nil {
		logs.Error(err.Error())
	}
	carsApply := auth.CarsApply{}
	qs.Filter("id", id).One(&carsApply)
	_, err = o.QueryTable("sys_cars").Filter("id", carsApply.Cars.Id).Update(orm.Params{
		"status": 1,
	})
	if err != nil {
		logs.Error(err.Error())
	}

	c.Redirect(beego.URLFor("CarsApplyController.MyApply"), 302)
}
