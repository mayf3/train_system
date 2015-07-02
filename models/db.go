package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"SBOJ/web/conf"
	. "SBOJ/web/modules/utils"
)

var (
	DATABASE        string = "mysql"
	CONNECTION      string = "iladmin:@tcp(localhost.localdomain:3306)/train"
	CONNECTION_FAIL string = "Can't connect to database"

	db *sql.DB
)

func InitDb() *sql.DB {
	var err error
	db, err = sql.Open(DATABASE, CONNECTION)
	RaisePanic(err, CONNECTION_FAIL)

	err = db.Ping()
	RaisePanic(err, CONNECTION_FAIL)
	return db
}

func GetDb() *sql.DB {
	return db
}

type DbAgent struct {
	from    string
	_select string
	where   string
	orderBy string
	join    []string

	err error
}

func (this *DbAgent) Clear() {
	this.from = ""
	this._select = ""
	this.where = ""
	this.orderBy = ""
	this.err = nil
	this.join = []string{}
}

func (this *DbAgent) From(table string) *DbAgent {
	this.from = table
	return this
}

func (this *DbAgent) Select(items []string) *DbAgent {
	this._select = ""
	for i := 0; i < len(items); i++ {
		if i > 0 {
			this._select += ", "
		}
		this._select += items[i]
	}
	return this
}

func (this *DbAgent) SelectAll() *DbAgent {
	this._select = "*"
	return this
}

func (this *DbAgent) Where(clause string, args ...interface{}) *DbAgent {
	size := len(args)
	cnt := 0

	this.where = ""
	for i := 0; i < len(clause); i++ {
		if EqualChar(clause[i], "?") {
			if cnt < size {
				this.where += ToSqlString(args[cnt])
			} else {
				//TODO: add error
				this.err = errors.New("Parameter's number no match")
				break
			}
			cnt++
		} else {
			this.where += string(clause[i])
		}
	}

	return this
}

func (this *DbAgent) OrderBy(order string) *DbAgent {
	this.orderBy = order
	return this
}

func (this *DbAgent) Join(table string) *DbAgent {
	this.join = append(this.join, table)
	return this
}

/**
 * generate insert sql like
 *
 * INSERT INTO table_name($, [$, [$...]]) VALUES($, [$, [$...]])
 */
func (this *DbAgent) Insert(info map[string]interface{}) (sql.Result, error) {
	err := this.checkError()
	if err != nil {
		return nil, err
	}

	fields := ""
	values := ""
	flag := false
	for key, val := range info {
		if flag {
			fields += ", "
			values += ", "
		}

		flag = true
		fields += key
		values += ToSqlString(val)
	}

	stat := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s);", this.from, fields, values)
	if config.Debug {
		fmt.Println(stat)
	}

	var result sql.Result
	result, err = db.Exec(stat)
	fmt.Println(db)

	if err != nil {
		return nil, err
	}
	return result, nil
}

/**
 * generate count sql like
 *
 * SELECT XXX, count(XXX) FROM T WHERE where_clause
 */
func (this *DbAgent) Count(args ...string) (*sql.Rows, error) {
	err := this.checkError()
	if err != nil {
		return nil, err
	}

	count := ""
	for i := 0; i < len(args); i++ {
		if i > 0 {
			count += ", "
		}
		count += fmt.Sprintf("count(%s)", args[i])
	}

	// If no args, then just count(*)
	if IsEmptyString(count) {
		count = "count(*)"
	}

	if !IsEmptyString(this._select) {
		this._select += ", "
	}
	this._select += count

	stat := fmt.Sprintf("SELECT %s FROM %s %s", this._select, this.from, this.getWhere())
	if config.Debug {
		fmt.Println(stat)
	}

	rows, err := db.Query(stat)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

/**
 * generate update sql like
 *
 * UPDATE T SET x=y, ... WHERE where_clause
 */
func (this *DbAgent) Update(args map[string]interface{}) (sql.Result, error) {
	err := this.checkError()
	if err != nil {
		return nil, err
	}

	setStat := ""
	flag := false
	for key, val := range args {
		if flag {
			setStat += ", "
		}
		setStat += fmt.Sprintf("%s=%s", key, ToSqlString(val))
		flag = true
	}

	stat := fmt.Sprintf("UPDATE %s SET %s %s", this.from, setStat, this.getWhere())

	if config.Debug {
		fmt.Println(stat)
	}

	var result sql.Result
	result, err = db.Exec(stat)
	if err != nil {
		return nil, err
	}
	return result, err
}

/**
 * genenrate delete sql like
 *
 * DELETE FROM T WHERE where_clause
 */
func (this *DbAgent) Delete() (sql.Result, error) {
	err := this.checkError()
	if err != nil {
		return nil, err
	}

	where_clause := this.getWhere()
	if len(where_clause) <= 0 {
		return nil, errors.New("No condition applied for delete")
	}

	stat := fmt.Sprintf("DELETE FROM %s %s", this.from, where_clause)
	if config.Debug {
		fmt.Println(stat)
	}

	var result sql.Result
	result, err = db.Exec(stat)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/**
 * generate query sql like
 *
 * SELECT $, [$...] FROM $ [ JOIN $...] WHERE ... ORDER BY ...
 *
 * TODO: add group_by
 */
func (this *DbAgent) Query() (*sql.Rows, error) {
	err := this.checkError()
	if err != nil {
		return nil, err
	}

	var rows *sql.Rows
	rows, err = db.Query(this.getQuery())
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (this *DbAgent) QueryRow() (*sql.Row, error) {
	err := this.checkError()
	if err != nil {
		return nil, err
	}

	row := db.QueryRow(this.getQuery())
	return row, nil
}

/**
 * Support for complex query
 */
func (this *DbAgent) RawQuery(stat string) (*sql.Rows, error) {
	rows, err := db.Query(stat)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (this *DbAgent) checkError() error {
	if this.err != nil {
		return this.err
	}

	if IsEmptyString(this.from) {
		return errors.New("No table has selected")
	}

	return nil
}

func (this *DbAgent) getWhere() string {
	if !IsEmptyString(this.where) {
		return "where " + this.where
	}
	return ""
}

func (this *DbAgent) getQuery() string {
	if len(this._select) <= 0 {
		this._select = "*"
	}
	stat := fmt.Sprintf("SELECT %s FROM %s ", this._select, this.from)

	join_clause := strings.Join(this.join, ", ")
	if len(join_clause) > 0 {
		stat += join_clause
	}

	where_clause := this.getWhere()
	if len(where_clause) > 0 {
		stat += where_clause
	}

	if len(this.orderBy) > 0 {
		stat += this.orderBy
	}
	return stat
}
