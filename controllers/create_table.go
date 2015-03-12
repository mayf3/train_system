package controllers

import (
	"github.com/astaxie/beego"
//	"train_system/models"
)

type table struct {
	title		string	`form:"title"`
	problem_number	string	`form:"number"`
	source		string	`form:"source"`
}

type CreateTableController struct {
	beego.Controller
}

func (c *CreateTableController) Get() {
	c.TplNames = "create_table.tpl"
}

func (c *CreateTableController) Post() {
	new_table := table{}
	if err := c.ParseForm(&new_table); err == nil {
		maps := map[string]string{}
		maps["Contest_name"] = new_table.title
		maps["Problem_number"] = new_table.problem_number
		maps["Source"] = new_table.source;
//		models.CreateTable(maps);
	}
}
