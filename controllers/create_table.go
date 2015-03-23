package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	. "train_system/models"
)

type CreateTableController struct {
	beego.Controller
}

func (c *CreateTableController) Get() {
	var (
		err   error
		tmp   string
		table Tables
	)
	tmp = c.Input().Get("table_id")
	table.Id, err = strconv.Atoi(tmp)
	if err == nil && table.Id != 0 {
		err = table.GetTableById(table.Id)
		c.Data["Init"] = table
	}
	c.TplNames = "create_table.tpl"
}

func (c *CreateTableController) Post() {
	var (
		err   error
		table Tables
	)
	table.Id, err = strconv.Atoi(c.GetString("table_id"))
	err = table.GetTableById(table.Id)
	table.ContestName = c.GetString("contest_name")
	table.ProblemNumber, err = strconv.Atoi(c.GetString("problem_number"))
	fmt.Println(table.ProblemNumber)
	fmt.Println(err)
	table.Source = c.GetString("source")
	if table.Id != 0 {
		table.Update()
	} else {
		table.Insert()
	}
	c.Redirect("/index", 302)
}
