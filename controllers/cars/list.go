package cars

import (
	"github.com/astaxie/beego"
)

type CarsController struct {
	beego.Controller
}

func (c *CarsController) List() {
	c.TplName = "cars/cars_list.html"
}
