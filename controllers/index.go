package controllers

import (
	"github.com/astaxie/beego"
	"train_system/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	maps, err := models.GetAllTable()
	if err == nil{
		this.Data["ok"]		= "true"
		this.Data["test"]	= maps[0]
	}
	this.TplNames = "index.tpl"
}
