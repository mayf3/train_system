package controllers

import (
	"github.com/astaxie/beego"
	"train_system/models"
	"fmt"
)

type IndexController struct {
	beego.Controller
}


func (this *IndexController) Get() {
	maps, err := models.GetAllTable()
	fmt.Println("len = ", len(maps))
	if err == nil{
		this.Data["Map"] = maps
	}
	this.TplNames = "index.tpl"
}
