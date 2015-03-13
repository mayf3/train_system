package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Tables struct {
	Id            int `orm:"pk;auto"`
	ContestName   string
	ProblemNumber int            `orm:"null"`
	Source        string         `orm:"null"`
	CreateTime    time.Time      `orm:"auto_now_add;type(datetime)"`
	Information   []*Information `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Tables))
}

func GetAllTable() (all_table []orm.Params, err error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("tables").All(&all_tableh)
	if num > 0 && err == nil {
		return all_table, err
	}
	return all_table, err
}

func GetSingleTable(table_id int) (single_table orm.Params, err error) {
	o := orm.NewOrm()
	err := o.QueryTable("tables").Filter("Id", table_id).One(&single_table)
	if err == nil {
		return single_table, err
	}
	return nil, err
}

func InsertTable(single_table orm.Params) error {
	o := orm.NewOrm()
	insert_table := new(Tables)
	insert_table.ContestName = single_table["contest_name"].(string)
	insert_table.ProblemNumber = single_table["problem_number"].(int)
	insert_table.Source = single_table["source"].(string)
	id, err := o.Insert(insert_table)
	if id > 0 {
	}
	return err
}

func EditTable(single_table orm.Params) error {
	o := orm.NewOrm()
	edit_table := Tables{Id: single_table["Id"].(int)}
	err := o.Read(&edit_table)
	if err == nil {
		edit_table.ContestName = single_table["contest_name"].(string)
		edit_table.ProblemNumber = single_table["problem_number"].(int)
		edit_table.Source = single_table["source"].(string)
		var num int64
		num, err = o.Update(&edit_table)
		if num > 0 {
		}
	}
	return err
}

func DeleteTable(table_id int) error {
	o := orm.NewOrm()
	num, err := o.Delete(&Tables{Id: table_id})
	if num > 0 {
	}
	return err
}
