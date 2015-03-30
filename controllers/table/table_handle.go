package table

import (
	"errors"
	"strings"

	"train_system/models"
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
func (this *TableController) generate() (models.Tables, [][]string) {
	var (
		_     error
		err   error
		hust  models.Hust
		table models.Tables
	)
	table.Id, _ = this.getTableId()
	_ = table.GetTableById(table.Id)
	err = hust.GetHustByTabel(table)
	if err != nil {
		return table, nil
	}
	return table, this.splitHustTable(hust)
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

// utils

func (this *TableController) getTableId() (table_id int, err error) {
	table_id, err = this.GetInt(":table_id")
	if err != nil {
		return 0, errors.New("no table id")
	}
	return table_id, nil
}

func (this *TableController) splitHustTable(hust models.Hust) [][]string {
	var (
		hust_table     []string
		all_hust_table [][]string
	)
	hust_table = strings.Split(hust.Context, string(rune(32)))
	for _, val := range hust_table {
		all_hust_table = append(all_hust_table, strings.Split(val, string(rune(9))))
	}
	for i := 1; i < len(all_hust_table); i++ {
		all_hust_table[i] = all_hust_table[i][0 : len(all_hust_table[i])-1]
	}
	return all_hust_table
}

func (this *TableController) checkTable() {
	var (
		err   error
		table models.Tables
	)
	table.Id, err = this.getTableId()
	if err != nil {
		this.Abort("Error table id")
	}
	err = table.GetTableById(table.Id)
	if err != nil {
		this.Abort("Error table id")
	}
}
