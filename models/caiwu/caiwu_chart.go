package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type CaiwuData struct {
	Id             int       `orm:"pk;auto"`
	CaiwuDate      string    `orm:"description(财务月份);size(20);column(caiwu_date)"`
	SalesVolume    float64   `orm:"column(sales_volume);digits(10);decimals(2);description(本月销售额)"`
	StudentIncress int       `orm:"column(student_incess);default(0);description(学员增加数)"`
	Django         int       `orm:"description(django课程卖出数量);column(django)"`
	VueDjango      int       `orm:"description(vue+django课程卖出数量)"`
	Celery         int       `orm:"description(celery课程卖出数量);column(celery)"`
	CreateDate     time.Time `orm:"auto_now_add;type(datetime)"`
}

func (c *CaiwuData) TableName() string {
	return "sys_caiwu_data"
}

func init() {
	orm.RegisterModel(new(CaiwuData))
}
