package controllers

import (
	"github.com/astaxie/beego"
)

type EditImformationController struct {
	beego.Controller
}

func (c *EditImformationController) Get() {
	c.TplNames = "edit_imformation.tpl"
}
