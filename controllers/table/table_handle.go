package table

import (

	"github.com/mayf3/train_system/models"
)

// getPreTable
func (this *TableController) getPreTable() models.Tables {
	var (
		err   error
		table models.Tables
	)
	table.Id, err = this.getTableId()
	if err != nil {
		return models.Tables{}
	}
	err = table.GetTableById(table.Id)
	if err != nil {
		return models.Tables{}
	}
	return table
}

func (this *TableController) setting() {
	var (
		err   error
		table models.Tables
	)
	table.Id, err = this.getTableId()
	if err == nil {
		err = table.GetTableById(table.Id)
		if err != nil {
			table.Id = 0
		}
	}
	table.ContestName = this.GetString("contest_name")
	table.ProblemNumber, err = this.GetInt("problem_number")
	table.Source = this.GetString("source")
	if table.Id != 0 {
		table.Update()
	} else {
		table.Insert()
	}
}

// must make sure table exists
func (this *TableController) del() {
	var (
		_     error
		table models.Tables
	)
	table.Id, _ = this.getTableId()
	_ = table.GetTableById(table.Id)
	table.Delete()
}

// must make sure table exists
//TODO Get && Post do same work, move it to tools/tool.go
func (this *TableController) generate() [][]string {
	var (
		err   error
		hust  models.Hust
		table models.Tables
	)
	table = this.getTable()
	err = hust.GetHustByTabel(table)
	if err != nil {
		return nil
	}
	return this.splitHustTable(hust)
}

// must make sure table exists
func (this *TableController) submitHust() {
	var (
		_     error
		err   error
		hust  models.Hust
		table models.Tables
	)
	table.Id, _ = this.getTableId()
	_ = table.GetTableById(table.Id)
	err = hust.GetHustByTabel(table)
	hust.Context = this.GetString("hust_table")
	hust.Table = &table
	if err != nil {
		hust.Insert()
	}
	hust.Update()
}
