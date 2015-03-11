package controllers

import (
	"github.com/astaxie/beego"
)

type CreateImformationController struct {
	beego.Controller
}

func (c *CreateImformationController) Get() {
	c.TplNames = "create_imformation.tpl"
}
