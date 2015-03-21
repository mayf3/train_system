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
	var(
		err error
		tmp string
		table Tables
		all_table []Tables
	)
	tmp = this.Input().Get("admin")
	if err == nil && tmp == "true"{
		this.Data["show"] = 1
	}
	all_table, err = table.GetAllTable()
	if err == nil {
		this.Data["Init"] = all_table
	}
	this.TplNames = "index.tpl"
}
