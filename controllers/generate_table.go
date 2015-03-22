package controllers

import (
	//	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	. "train_system/models"
	. "train_system/tools"
)

type GenerateTableController struct {
	beego.Controller
}

//TODO Get && Post do same work, move it to tools/tool.go
func (this *GenerateTableController) Get() {
	var (
		_              error
		err            error
		tmp            string
		hust_table     []string
		all_hust_table [][]string
		hust           Hust
		table          Tables
	)
	tmp = this.Input().Get("table_id")
	table.Id, _ = strconv.Atoi(tmp)
	_ = table.GetTableById(table.Id)
	this.Data["title"] = table.ContestName
	this.Data["source"] = table.Source
	this.Data["date"] = table.CreateTime
	this.Data["problem_list"] = GenerateProblemList(table.Id)
	this.Data["total_status"] = GenerateTable(table.Id)
	err = hust.GetHustByTabel(table)
	if err == nil {
		hust_table = strings.Split(hust.Context, string(rune(32)))
		for _, val := range hust_table {
			all_hust_table = append(all_hust_table, strings.Split(val, string(rune(9))))
		}
		for i := 1; i < len(all_hust_table); i++ {
			all_hust_table[i] = all_hust_table[i][0 : len(all_hust_table[i])-1]
		}
		this.Data["hust_table"] = all_hust_table
	}
	this.TplNames = "generate_table.tpl"
}

func (this *GenerateTableController) Post() {
	var (
		_              error
		err            error
		tmp            string
		hust_table     []string
		all_hust_table [][]string
		hust           Hust
		table          Tables
	)
	tmp = this.Input().Get("table_id")
	hust_table = strings.Split(this.GetString("hust_table"), string(rune(32)))
	for _, val := range hust_table {
		all_hust_table = append(all_hust_table, strings.Split(val, string(rune(9))))
	}
	for i := 1; i < len(all_hust_table); i++ {
		all_hust_table[i] = all_hust_table[i][0 : len(all_hust_table[i])-1]
	}
	table.Id, _ = strconv.Atoi(tmp)
	_ = table.GetTableById(table.Id)
	err = hust.GetHustByTabel(table)
	if err == nil {
		hust.Context = this.GetString("hust_table")
		hust.Table = &table
		hust.Update()
	} else {
		hust.Context = this.GetString("hust_table")
		hust.Table = &table
		hust.Insert()
	}
	this.Data["title"] = table.ContestName
	this.Data["source"] = table.Source
	this.Data["date"] = table.CreateTime
	this.Data["problem_list"] = GenerateProblemList(table.Id)
	this.Data["total_status"] = GenerateTable(table.Id)
	this.Data["hust_table"] = all_hust_table
	this.TplNames = "generate_table.tpl"
}
