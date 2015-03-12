package controllers

import (
	"github.com/astaxie/beego"
	"train_system/models"
	"fmt"
)


type CreateTableController struct {
	beego.Controller
}

func (c *CreateTableController) Get() {
	c.TplNames = "create_table.tpl"
}

func (c *CreateTableController) Post() {
	title := c.GetString("title")
	number := c.GetString("number")
	source := c.GetString("source")
	maps := map[string]string{}
	maps["contest_name"] = title
	maps["problem_number"] = number
	maps["source"] = source
	fmt.Println(maps);
	models.InsertTable(maps)
	c.Redirect("/index", 302)
}
