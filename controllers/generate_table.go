package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	. "train_system/models"
)

type GenerateTableController struct {
	beego.Controller
}

func (c *GenerateTableController) Get() {
	var table Tables
	table_id := c.Input().Get("table_id")
	id, err := strconv.Atoi(table_id)
	err = table.GetTableById(id)
	if err == nil {
		fmt.Println(table)
		c.Data["title"] = table.ContestName
		c.Data["source"] = table.Source
		c.Data["date"] = table.CreateTime
	}
	c.TplNames = "generate_table.tpl"
}
