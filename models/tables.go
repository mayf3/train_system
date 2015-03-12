package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Tables struct{
	Id				int			`orm:"pk;auto"`
	ContestName		string
	ProblemNumber	int			`orm:"null"`
	Source			string		`orm:"null"`
	CreateTime		time.Time	`orm:"auto_now_add;type(datetime)"`
	Information		[]*Information	`orm:"reverse(many)"`
}

func init(){
	orm.RegisterModel(new(Tables))
}

func GetAllTable() (maps []orm.Params, err error){
	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM tables").Values(&maps)
	if num > 0 && err == nil {
		return maps, err
	}
	return maps, err
}

func GetSingleTable(table_id int) (orm.Params, error){
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM tables WHERE Id=?", table_id).Values(&maps)
	if num > 0 && err == nil{
		return maps[0], err
	}
	return nil, err
}

func InsertTable(single_table map[string]string) error{
	o := orm.NewOrm()
	res, err := o.Raw("INSERT INTO tables(Contest_Name, Problem_Number, Source) VALUES (?, ?, ?)", single_table["contest_name"], single_table["problem_number"], single_table["source"]).Exec()
	if res == nil {}
	return err
}

func EditTable(single_table orm.Params) error{
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE tables(ContestName, ProblemNumber, Source) VALUES (?, ?, ?) WHERE Id = ?", single_table["contest_name"], single_table["problem_number"], single_table["source"], single_table["id"]).Exec()
	if res == nil {}
	return err
}

func DeleteTable(table_id int) error{
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM tables WHERE Id = ?", table_id).Exec()
	if res == nil {}
	return err
}
