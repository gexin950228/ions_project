package auth

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Auth struct {
	Id         int       `orm:"pk;auto"`
	AuthName   string    `orm:"description(权限名称);size(64)"`
	UrlFor     string    `orm:"description(url反转);size(255)"`
	Pid        int       `orm:"description(父节点id)"`
	Desc       string    `orm:"description(描述);size(255)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`
	IsActive   int       `orm:"column(is_active);description(1启用，0停用)"`
	IsDelete   int       `orm:"columns(is_delete);description(1删除，0未删除)"`
	Weight     int       `orm:"column(weight);description(权重，数值越大，权重越大)"`
	PName      string    `orm:"size(50);column(p_name)"`
}

type Role struct {
	Id         int       `orm:"pk;auto"`
	RoleName   string    `orm:"size(64)"`
	IsActive   int       `orm:"column(is_active)"`
	Desc       string    `orm:"size(255)"`
	IsDelete   int       `orm:"column(is_delete)"`
	CreateTime time.Time `orm:"type(datetime);auto_now;description(创建时间)"`
	Auth       []*Auth   `orm:"rel(m2m)"`
	User       []*User   `orm:"rel(m2m)"`
}

type User struct {
	Id            int              `orm:"pk;auto"`
	CardId        string           `orm:"size(64);column(card_id);description(员工工号);column(card_id)"`
	UserName      string           `orm:"unique;column(user_name);size(64);description(用户名);column(user_name)"`
	Password      string           `orm:"size(32);description(密码);column(password)"`
	Age           int              `orm:"null;description(年龄);column(age)"`
	Gender        int              `orm:"null;description(1:男,2:女,3:未知);column(gender)"`
	Phone         int64            `orm:"null;description(电话号码);column(phone)"`
	Addr          string           `orm:"null;size(255);description(地址);column(addr)"`
	IsActive      int              `orm:"description(1启用，0停用);default(1);"`
	IsDelete      int              `orm:"description(1删除，0未删除);default(0)"`
	CreateTime    time.Time        `orm:"auto_now;type(datetime);description(创建时间);null"`
	Role          []*Role          `orm:"reverse(many)"`
	CarsApply     []*CarsApply     `orm:"reverse(many)"`
	MessageNotify []*MessageNotify `orm:"reverse(many)"`
}

type CarBrand struct {
	Id         int       `orm:"pk;auto"`
	Name       string    `orm:"description(品牌名称);size(64)"`
	Desc       string    `orm:"description(品牌描述);size(255)"`
	Cars       []*Cars   `orm:"reverse(many)"`
	IsActive   int       `orm:"default(1);description(启用:1,停用:0)"`
	IsDelete   int       `orm:"default(0);description(删除:1,未删除:0)"`
	CreateTime time.Time `orm:"description(创建时间);type(datetime);auto_now"`
}

// Cars 车辆列表
type Cars struct {
	Id         int          `orm:"pk;auto"`
	Name       string       `orm:"description(车辆名称);size(64)"`
	CarBrand   *CarBrand    `orm:"rel(fk);description(车辆品牌外键)"`
	CarsApply  []*CarsApply `orm:"reverse(many)"`
	Status     int          `orm:"description(1:可借,2:不可借);default(1)"`
	IsActive   int          `orm:"default(1);description(启用:1,停用:0)"`
	IsDelete   int          `orm:"default(0);description(删除:1,未删除:0)"`
	CreateTime time.Time    `orm:"description(创建时间);type(datetime);auto_now"`
}

// CarsApply 车辆申请、审核
type CarsApply struct {
	Id           int       `orm:"pk;auto"`
	User         *User     `orm:"rel(fk)"`
	Cars         *Cars     `orm:"rel(fk)"`
	Reason       string    `orm:"description(申请理由);size(255)"`
	Destination  string    `orm:"description(目的地);size(64)"`
	ReturnDate   time.Time `orm:"type(date);auto_now;description(归还日期)"`
	ReturnStatus int       `orm:"descrption(1:已归还，0：未归还);default(0)"`
	AuditStatus  int       `orm:"description(1:同意，2:未同意，3:未审批);default(3)"`
	AuditOption  string    `orm:"description(审批意见);size(255)"`
	IsActive     int       `orm:"default(1);description(启用:1,停用:0)"`
	IsDelete     int       `orm:"default(0);description(删除:1,未删除:0)"`
	CreateTime   time.Time `orm:"description(创建时间);type(datetime);auto_now"`
	NotifyTag    int       `orm:"description(1:已发送通知，0：未发送通知);default(0)"`
}

// MessageNotify 消息通知
type MessageNotify struct {
	Id      int    `orm:"pk;auto"`
	Flag    int    `orm:"description(1:车辆逾期，2:所有通知);default(1)"`
	Title   string `orm:"size(64);description(消息标题)"`
	Content string `orm:"type(text);description(消息内容)"`
	User    *User  `orm:"rel(fk);desription(用户外键)"`
	ReadTag int    `orm:"description(1:已读，0:未读)"`
}

func (u *User) TableName() string {
	return "sys_user"
}

func (r *Auth) TableName() string {
	return "sys_auth"

}

func (r *Role) TableName() string {
	return "sys_role"

}

func (u *CarBrand) TableName() string {
	return "sys_cars_brand"
}

func (u *Cars) TableName() string {
	return "sys_cars"
}

func (u *CarsApply) TableName() string {
	return "sys_cars_apply"
}

func (u *MessageNotify) TableName() string {
	return "sys_message_notify"
}

func init() {
	username := beego.AppConfig.String("username")
	pwd := beego.AppConfig.String("pwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	db := beego.AppConfig.String("db")

	// username:pwd@tcp(ip:port)/db?charset=utf8&loc=Local
	dataSource := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8&loc=Local"
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = true
	if err != nil {
		logs.Error(err)
	}
	err = orm.RegisterDataBase("default", "mysql", dataSource)
	if err != nil {
		logs.Error(err)
	}

	ret := fmt.Sprintf("host:%s|port:%s|db:%s", host, port, db)
	logs.Info(ret)
	orm.RegisterModel(
		new(Auth),
		new(Role),
		new(User),
		new(CarBrand),
		new(Cars),
		new(CarsApply),
		new(MessageNotify),
		new(Tree),
	)
}
