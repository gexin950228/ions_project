package my_center

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type SalarySlip struct {
	Id                 int       `orm:"pk;auto"`
	CardId             string    `orm:"column(card_id);size(64);description(员工工号)"`
	BasePay            float64   `orm:"column(base_pay);digits(12);decimals(2);description(基本工资)"`
	WorkingDays        float64   `orm:"column(working_days);digits(3);description(工作天数);decimals(1)"`
	DaysOff            float64   `orm:"description(请假天数);digits(3);decimals(1)"`
	DaysOffNo          float64   `orm:"description(调休天数);digits(3);decimals(1)"`
	Reward             float64   `orm:"description(奖金);column(reward);digits(11);decimals(2)"`
	RentSubsidy        float64   `orm:"description(租房补贴);column(rent_subsidy);digits(11);decimals(2)"`
	TransSubsidy       float64   `orm:"description(交通补贴);column(trans_subsidy);digits(11);decimals(2)"`
	SocialSecurity     float64   `orm:"description(社保);column(social_security);digits(11);decimals(2)"`
	HouseProvidentFund float64   `orm:"column(house_provident_fund);description(住房公积金);digits(11);decimals(2)"`
	PersonalPncomeTax  float64   `orm:"column(personal_pncome_tax);description(个税);digits(11);decimals(2)"`
	Fine               float64   `orm:"column(fine);description(罚金);digits(11);decimals(2)"`
	NetSalary          float64   `orm:"description(实发工资);column(net_salary);digits(11);decimals(2)"`
	PayDate            string    `orm:"column(pay_date);description(工资月份);size(32)"`
	CreateTime         time.Time `orm:"column(create_time);auto_now;type(datetime);description(创建时间)"`
}

func (s SalarySlip) TableName() string {
	return "sys_salary_slip"
}

func init() {
	orm.RegisterModel(new(SalarySlip))
}
