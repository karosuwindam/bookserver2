package table

import (
	"database/sql"
	"errors"
	"reflect"
)

type tablename string

const (
	Booknames tablename = "booknames"
	Filelists tablename = "filelists"
	Copyfile  tablename = "copyfile"
)

func dballread(sqltype string, rows *sql.Rows, t_name tablename) (interface{}, error) {

	switch t_name {
	case Booknames:
		return booknames_Read(sqltype, rows)
	case Filelists:
		return filelists_Read(sqltype, rows)
	case Copyfile:
		return copyfile_Read(sqltype, rows)
	default:
	}
	return 0, errors.New("Do note user db table")
}

func Read_Sqlite3(rows *sql.Rows, stdata ...any) ([]any, error) {
	var loaddata []interface{}
	rt := reflect.TypeOf(stdata[0])
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		switch f.Type.Kind() {
		case reflect.Int:
			i := int64(0)
			loaddata = append(loaddata, &i)
		case reflect.String:
			str := string("")
			loaddata = append(loaddata, &str)
		case reflect.Struct:
			str := string("")
			loaddata = append(loaddata, &str)
		default:
		}
	}
	err := rows.Scan(loaddata...)
	var output []interface{}
	for _, p := range loaddata {
		switch p.(type) {
		case *int64:
			tmp := p.(*int64)
			output = append(output, int(*tmp))
		case *string:
			tmp := p.(*string)
			output = append(output, *tmp)
		}
	}

	return output, err
}
