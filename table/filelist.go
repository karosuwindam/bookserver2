package table

import (
	"database/sql"
	"errors"
	"reflect"
	"strconv"
	"time"
)

//特定ファイルとリンクするzipとpdfのファイル名やTag
//主に、表の画面で使用される

type filelists struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Pdfpass    string    `json:"pdfpass"`
	Zippass    string    `json:"zippass`
	Tag        string    `json:"tag"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func convert_filelist(v ...any) (filelists, error) {
	var output filelists
	var err error = nil
	if len(v) != reflect.TypeOf(filelists{}).NumField() {
		return filelists{}, errors.New("Input data count err " + strconv.Itoa(len(v)))
	}
	output.Id = v[0].(int)
	output.Name = v[1].(string)
	output.Pdfpass = v[2].(string)
	output.Zippass = v[3].(string)
	output.Tag = v[4].(string)
	ctime := v[5].(string)
	utime := v[6].(string)
	output.Created_at, err = timeconvert(ctime)
	output.Updated_at, err = timeconvert(utime)

	return output, err
}

func filelists_Read(sqltype string, rows *sql.Rows) (filelists, error) {
	switch sqltype {
	case "sqlite3":
		return filelists_Read_sqlite3(rows)
	case "mysql":
		return filelists_Read_mysql(rows)
	default:
	}
	return filelists{}, errors.New("Don't user sql type for" + sqltype)
}

func filelists_Read_sqlite3(rows *sql.Rows) (filelists, error) {
	data := filelists{}
	tmp, err := Read_Sqlite3(rows, data)
	if err != nil {

	}
	output, err := convert_filelist(tmp...)

	return output, nil
}
func filelists_Read_mysql(rows *sql.Rows) (filelists, error) {
	return filelists{}, nil
}
