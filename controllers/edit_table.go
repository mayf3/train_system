package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	. "train_system/models"
)

type EditTableController struct {
	beego.Controller
}

func (c *EditTableController) Get() {
	var table Tables
	table_id := c.Input().Get("table_id")
	c.Data["table_id"] = table_id
	id, err := strconv.Atoi(table_id)
	err = table.GetTableById(id)
	if err == nil{
		fmt.Println("aa")
		var problem_list []string
		for i := 0; i < table.ProblemNumber; i++ {
			problem_list = append(problem_list, string(rune(i + 65)))
		}
		fmt.Println(problem_list)
		c.Data["problem_list"] = problem_list
	}
	c.TplNames = "edit_table.tpl"
}
