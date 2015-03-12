package controllers

import (
	"github.com/astaxie/beego"
)

type GenerateTableController struct {
	beego.Controller
}

func (c *GenerateTableController) Get() {
	table_id := c.Input().Get("table_id");
	c.Data["title"] = table_id
	maps, err := models.GetSingleTable(table_id);
	//if err == nil {
	//	c.Data["title"] = maps["title"];
	//	c.Data["source"] = maps["source"];
	//	c.Data["date"] = maps["date"];
	//}
	c.TplNames = "generate_table.tpl"
}
