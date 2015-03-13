package controllers

import (
	"github.com/astaxie/beego"
	//	"train_system/models"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	info := []map[string]string{}
	//info = models.GetTables();
	info = append(info, map[string]string{"date": "1", "title": "first", "source": "fs"})
	info = append(info, map[string]string{"date": "2", "title": "second", "source": "ss"})
	c.Data["Map"] = info
	c.TplNames = "index.tpl"
}
