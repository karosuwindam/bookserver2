package table

import (
	"database/sql"
	"errors"
)

type tablename string

const (
	Booknames tablename = "booknames"
)

func dbread(sqltype string, rows *sql.Rows, t_name string) (interface{}, error) {

	switch t_name {
	case "booknames":
		return booknames_Read(sqltype, rows)
	default:
	}
	return 0, errors.New("Do note user db table")
}
