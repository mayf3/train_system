package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

// TODO make a tool.go file for other function like change_string_to_int

type Tables struct {
	Id            int `orm:"pk;auto"`
	ContestName   string
	ProblemNumber int
	Source        string
	CreateTime    time.Time      `orm:"auto_now_add;type(datetime)"`
	Information   []*Information `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Tables))
}

func GetAllTable() ([]Tables, error) {
	o := orm.NewOrm()
	var all_table []Tables
	_, err := o.QueryTable("tables").All(&all_table)
	return all_table, err
}

func GetSingleTable(table_id int) (Tables, error) {
	o := orm.NewOrm()
	var single_table Tables
	err := o.QueryTable("tables").Filter("Id", table_id).One(&single_table)
	return single_table, err
}

func InsertTable(single_table map[string]string) error {
	var _ error
	o := orm.NewOrm()
	insert_table := new(Tables)
	insert_table.ContestName = single_table["contest_name"]
	insert_table.ProblemNumber, _ = strconv.Atoi(single_table["problem_number"])
	insert_table.Source = single_table["source"]
	_, err := o.Insert(insert_table)
	return err
}

func EditTable(single_table map[string]string) error {
	var _ error
	var num int
	num, _ = strconv.Atoi(single_table["id"])
	o := orm.NewOrm()
	edit_table := Tables{Id: num}
	err := o.Read(&edit_table)
	if err == nil {
		edit_table.ContestName = single_table["contest_name"]
		edit_table.ProblemNumber, _ = strconv.Atoi(single_table["problem_number"])
		edit_table.Source = single_table["source"]
		var _ int64
		_, err = o.Update(&edit_table)
	}
	return err
}

func DeleteTable(table_id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Tables{Id: table_id})
	return err
}
