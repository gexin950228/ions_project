package news

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Category struct {
	Id         int       `orm:"pk;auto"`
	Name       string    `orm:"size(64);description(分类名称)"`
	Desc       string    `orm:"size(255);description(分类描述)"`
	IsActive   int       `orm:"default(1);description(启用:1;停用:0)"`
	IsDelete   int       `orm:"default(0);description(删除:1；未删除:0)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`
	News       []*News   `orm:"reverse(many)"`
}

type News struct {
	Id         int       `orm:"pk;auto"`
	Title      string    `orm:"size(64);description(新闻标题)"`
	Content    string    `orm:"size(255);description(新闻内容);type(text)"`
	IsActive   int       `orm:"default(1);description(启用:1,停用:0)"`
	IsDelete   int       `orm:"default(0);description(删除:1,未删除:0)"`
	Category   *Category `orm:"rel(fk)"`
	CreateTime time.Time `orm:"description(创建时间);type(datetime);auto_now"`
}

func (c *Category) TableName() string {
	return "sys_category"
}

func (n *News) TableName() string {
	return "sys_news"
}

func init() {
	orm.RegisterModel(new(Category), new(News))
}
