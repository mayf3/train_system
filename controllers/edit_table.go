package controllers

import (
	"strconv"
	"github.com/astaxie/beego"
	. "train_system/models"
	. "train_system/tools"
)

type EditTableController struct {
	beego.Controller
}

func (c *EditTableController) Get() {
	var (
		_ error
		tmp string
		table Tables
	)
	tmp = c.Input().Get("table_id")
	c.Data["table_id"] = tmp
	table.Id, _ = strconv.Atoi(tmp)
	_ = table.GetTableById(table.Id)
	c.Data["problem_list"] = GenerateProblemList(table.Id)
	c.Data["total_status"] = GenerateTable(table.Id)
	c.TplNames = "edit_table.tpl"
}
