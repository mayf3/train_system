package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	. "train_system/models"
	. "train_system/tools"
)

type GenerateTableController struct {
	beego.Controller
}

func (c *GenerateTableController) Get() {
	var(
		_ error
		tmp string
		table Tables
	)
	tmp = c.Input().Get("table_id")
	table.Id, _ = strconv.Atoi(tmp)
	_ = table.GetTableById(table.Id)
	c.Data["title"] = table.ContestName
	c.Data["source"] = table.Source
	c.Data["date"] = table.CreateTime
	c.Data["problem_list"] = GenerateProblemList(table.Id)
	c.Data["total_status"] = GenerateTable(table.Id)
	c.TplNames = "generate_table.tpl"
}
