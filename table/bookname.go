package table

import (
	"database/sql"
	"errors"
	"time"
)

type booknames struct {
	Id         int       `json:"id" db:"id" type:"int"`
	Name       string    `json:"name" db:"name" type:"string"`
	Title      string    `json:"title" db:"title" type:"string"`
	Writer     string    `json:"writer db:"writer" type:"string"`
	Brand      string    `json:"brand" db:"brand" type:"string"`
	Booktype   string    `json:"booktype" db:"booktype" type:"string"`
	Ext        string    `json:"ext" db:"ext" type:"string"`
	Created_at time.Time `json:"created_at" db:"created_at" type:"time"`
	Updated_at time.Time `json:"updated_at" db:"updated_at" type:"time`
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
	var err error
	tmp := booknames{}
	var layout1 = "2006-01-02T15:04:05.999999Z"
	var layout2 = "2006-01-02 15:04:05.999999999"
	c_time_tmp := ""
	u_time_tmp := ""
	if err := rows.Scan(&tmp.Id, &tmp.Name, &tmp.Title, &tmp.Writer, &tmp.Brand, &tmp.Booktype, &tmp.Ext, &c_time_tmp, &u_time_tmp); err != nil {
		return tmp, err
	}
	if tmp.Created_at, err = time.Parse(layout1, c_time_tmp); err != nil {
		if tmp.Created_at, err = time.Parse(layout2, c_time_tmp); err != nil {
			return tmp, err
		}
	}
	if tmp.Updated_at, err = time.Parse(layout1, u_time_tmp); err != nil {
		if tmp.Updated_at, err = time.Parse(layout2, u_time_tmp); err != nil {
			return tmp, err
		}
	}
	return tmp, nil
}

func booknames_Read_mysql(rows *sql.Rows) (booknames, error) {
	tmp := booknames{}
	return tmp, nil
}
