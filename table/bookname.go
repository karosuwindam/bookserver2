package table

import (
	"database/sql"
	"errors"
	"reflect"
	"strconv"
	"time"
)

//読み書き用のベースになるデータベース

type booknames struct {
	Id         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Title      string    `json:"title" db:"title"`
	Writer     string    `json:"writer" db:"writer"`
	Brand      string    `json:"brand" db:"brand"`
	Booktype   string    `json:"booktype" db:"booktype"`
	Ext        string    `json:"ext" db:"ext"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
}

func convert_booknames(v ...any) (booknames, error) {
	var output booknames
	var err error = nil
	if len(v) != reflect.TypeOf(booknames{}).NumField() {
		return booknames{}, errors.New("Input data count err " + strconv.Itoa(len(v)))
	}
	i := 0
	output.Id = v[i].(int)
	i++
	output.Name = v[i].(string)
	i++
	output.Title = v[i].(string)
	i++
	output.Writer = v[i].(string)
	i++
	output.Brand = v[i].(string)
	i++
	output.Booktype = v[i].(string)
	i++
	output.Ext = v[i].(string)
	i++
	output.Created_at = v[i].(time.Time)
	i++
	output.Updated_at = v[i].(time.Time)

	return output, err
}

func booknames_table_used(v ...any) bool {
	return false
}

func booknames_Update(sqltype string, rows *sql.Rows, id int, v ...any) {

}

func booknames_Insart(sqltype string, rows *sql.Rows) {

}

func booknames_Read(sqltype string, rows *sql.Rows) (booknames, error) {
	switch sqltype {
	case "sqlite3":
		return booknames_Read_sqlite3(rows)
	case "mysql":
		return booknames_Read_mysql(rows)
	default:
	}
	return booknames{}, errors.New("Don't user sql type for" + sqltype)
}

func booknames_Read_sqlite3(rows *sql.Rows) (booknames, error) {
	data := booknames{}
	tmp, err := Read_Sqlite3(rows, data)
	if err != nil {
		return booknames{}, err
	}
	output, err := convert_booknames(tmp...)

	return output, nil
}

func booknames_Read_mysql(rows *sql.Rows) (booknames, error) {
	tmp := booknames{}
	return tmp, nil
}
