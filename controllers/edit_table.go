package controllers

import (
	"github.com/astaxie/beego"
)

type EditTableController struct {
	beego.Controller
}

func (c *EditTableController) Get() {
	c.TplNames = "edit_table.tpl"
}
