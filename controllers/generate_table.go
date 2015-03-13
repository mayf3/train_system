package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"train_system/models"
)

type GenerateTableController struct {
	beego.Controller
}

func (c *GenerateTableController) Get() {
	table_id := c.Input().Get("table_id")
	id, err := strconv.Atoi(table_id)
	maps, err := models.GetSingleTable(id)
	if err == nil {
		fmt.Println("aa")
		c.Data["title"] = maps.ContestName
		c.Data["source"] = maps.Source
		//c.Data["date"] = maps.CreateTime
	}
	c.TplNames = "generate_table.tpl"
}
