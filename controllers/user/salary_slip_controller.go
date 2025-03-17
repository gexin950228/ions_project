package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_project/models/auth"
	"ions_project/models/my_center"
	"time"
)

type SalarySlipController struct {
	beego.Controller
}

func (c *SalarySlipController) Get() {
	month := c.GetString("month")
	if month == "" {
		month = time.Now().Format("2006-01")
	}
	fmt.Printf("month:%s\n", month)
	id := c.GetSession("id")
	if id == "" {
		id = "1"
	}
	o := orm.NewOrm()
	user := auth.User{}
	o.QueryTable("sys_user").Filter("id", id).One(&user)
	fmt.Printf("user card_id:%#v, %T\n", user.CardId, user.CardId)
	cardId := user.CardId
	salarySlip := my_center.SalarySlip{}
	err := o.QueryTable("sys_salary_slip").Filter("card_id", cardId).Filter("pay_date", month).One(&salarySlip)
	if err != nil {
		logs.Error(fmt.Sprintf("查询当月工资出错，错误信息: %s", err.Error()))
	}
	fmt.Println(salarySlip)
	c.Data["salary_slip"] = salarySlip
	c.TplName = "user/salary_slip_detail.html"
}

func (c *SalarySlipController) Detail() {
	id := c.GetSession("id")
	o := orm.NewOrm()
	salary_slip := my_center.SalarySlip{}
	err := o.QueryTable("sys_salary_slip").Filter("id", id).One(&salary_slip)
	if err != nil {
		logs.Error(err.Error())
	}

	c.Data["salary_slip"] = salary_slip
	c.TplName = "user/salary_slip_detail.html"
}
