package controllers

import (
	"github.com/astaxie/beego"
//	"strconv"
//	"train_system/models"
)

type ActionController struct {
	beego.Controller
}

/*
func (this *ActionController) Get() {
	action := this.GetString("action")
	if action == "DeleteTable" {
		table_id := this.Input().Get("id")
		id, _ := strconv.Atoi(table_id)
		DeleteTable(id)
		this.Redirect("/index", 302)
	}
}

func DeleteTable(id int) {
	models.DeleteTable(id)
}
*/
