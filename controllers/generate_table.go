package controllers

import (
	"github.com/astaxie/beego"
	"train_system/models"
	"strconv"
)

type GenerateTableController struct {
	beego.Controller
}

func (c *GenerateTableController) Get() {
	table_id := c.Input().Get("table_id")
	id, err := strconv.Atoi(table_id)
	maps, err := models.GetSingleTable(id)
	if err == nil {
		c.Data["title"] = maps["contest_name"]
		c.Data["source"] = maps["source"]
		c.Data["date"] = maps["create_time"]
	}
	c.TplNames = "generate_table.tpl"
}
