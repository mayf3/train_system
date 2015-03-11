package controllers

import (
	"github.com/astaxie/beego"
)

type CreateTableController struct {
	beego.Controller
}

func (c *CreateTableController) Get() {
	c.TplNames = "create_table.tpl"
}
