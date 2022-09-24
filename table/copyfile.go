package table

import (
	"database/sql"
	"errors"
	"reflect"
	"strconv"
	"time"
)

// NextCloud連携用のSamba共有しているフォルダについて置いてあるFileの記録
type copyfile struct {
	Id         int       `json:"id" db:"id"`
	Zippass    string    `json:"zippass" db:"zippass"`
	Filesize   int       `json:"filesize" db:"filesize"`
	Copyflag   int       `json:"copyflag" db:"copyflag"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
}

func convert_copyfile(v ...any) (copyfile, error) {
	var output copyfile
	var err error = nil
	if len(v) != reflect.TypeOf(copyfile{}).NumField() {
		return copyfile{}, errors.New("Input data count err " + strconv.Itoa(len(v)))
	}
	i := 0
	output.Id = v[i].(int)
	i++
	output.Zippass = v[i].(string)
	i++
	output.Filesize = v[i].(int)
	i++
	output.Copyflag = v[i].(int)
	i++
	output.Created_at = v[i].(time.Time)
	i++
	output.Updated_at = v[i].(time.Time)

	return output, err
}

func copyfile_table_used(v ...any) bool {
	return false
}

func copyfile_Update(sqltype string, rows *sql.Rows, id int, v ...any) {

}

func copyfile_Insart(sqltype string, rows *sql.Rows) {

}

func copyfile_Read(sqltype string, rows *sql.Rows) (copyfile, error) {
	switch sqltype {
	case "sqlite3":
		return copyfile_Read_sqlite3(rows)
	case "mysql":
		return copyfile_Read_mysql(rows)
	default:
	}
	return copyfile{}, errors.New("Don't user sql type for" + sqltype)
}

func copyfile_Read_sqlite3(rows *sql.Rows) (copyfile, error) {
	data := copyfile{}
	tmp, err := Read_Sqlite3(rows, data)
	if err != nil {
		return copyfile{}, err
	}
	output, err := convert_copyfile(tmp...)

	return output, nil
}

func copyfile_Read_mysql(rows *sql.Rows) (copyfile, error) {
	tmp := copyfile{}
	return tmp, nil
}
