package controllers

import (
//	"fmt"
	"github.com/astaxie/beego"
	. "train_system/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	var tmp_table Tables
	maps, err := tmp_table.GetAllTable()
	if err == nil {
		this.Data["Map"] = maps
	}
	this.TplNames = "index.tpl"
}
