package controllers

import (
	//"fmt"
	"strconv"
	"github.com/astaxie/beego"
	. "train_system/models"
)

type CreateTableController struct {
	beego.Controller
}

func (c *CreateTableController) Get() {
	var tmp_table Tables
	table_id := c.Input().Get("table_id")
	id, err := strconv.Atoi(table_id)
	if err == nil && id != 0 {
		err = tmp_table.GetTableById(id)
		c.Data["Init"] = tmp_table
	}
	c.TplNames = "create_table.tpl"
}

func (c *CreateTableController) Post() {
	var table Tables
	var err	error
	table.Id, err = strconv.Atoi(c.GetString("table_id"))
	err = table.GetTableById(table.Id)
	table.ContestName = c.GetString("contest_name")
	table.ProblemNumber, err = strconv.Atoi(c.GetString("problem_number"))
	table.Source = c.GetString("source")
	if err == nil && table.Id != 0 {
		table.Update()
	} else {
		table.Insert()
	}
	c.Redirect("/index", 302)
}
