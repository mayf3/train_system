package controllers

import (
	"github.com/astaxie/beego"
)

type GenerateTableController struct {
	beego.Controller
}

func (c *GenerateTableController) Get() {
	c.TplNames = "generate_table.tpl"
}
