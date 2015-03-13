package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"train_system/models"
)

type CreateTableController struct {
	beego.Controller
}

func (c *CreateTableController) Get() {
	table_id := c.Input().Get("table_id")
	id, err := strconv.Atoi(table_id)
	if err == nil && id != 0 {
		maps, _ := models.GetSingleTable(id)
		c.Data["Post"] = maps
	}
	c.TplNames = "create_table.tpl"
}

func (c *CreateTableController) Post() {
	table_id := c.GetString("table_id")
	id, err := strconv.Atoi(table_id)
	title := c.GetString("title")
	number := c.GetString("number")
	source := c.GetString("source")
	maps := map[string]string{}
	maps["contest_name"] = title
	maps["problem_number"] = number
	maps["source"] = source
	if err == nil && id != 0 {
		maps["id"] = table_id
		models.EditTable(maps)
	} else {
		models.InsertTable(maps)
	}
	fmt.Println(maps)
	c.Redirect("/index", 302)
}
