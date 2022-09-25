package table

import (
	"bookserver/message"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"time"
)

type Tablename string

type createTableS struct {
	name Tablename
	cmd  string
	st   interface{}
}

const (
	Booknames Tablename = "booknames"
	Filelists Tablename = "filelists"
	Copyfile  Tablename = "copyfile"
)

func tablecreate(sqltype string, sql *sql.DB) error {
	var createTableList []createTableS
	for tablename, stype := range tablemap {
		createTableList = append(createTableList, createTableS{name: tablename, st: stype})

	}

	for i := 0; i < len(createTableList); i++ {
		err := createTableList[i].createTableCmd(sqltype)
		if err != nil {
			return err
		}
	}
	createdTableList, err := tableList(sqltype, sql)
	if err != nil {
		return err
	}
	for _, list := range createTableList {
		create_flag := true
		for _, s := range createdTableList {
			if s == string(list.name) {
				create_flag = false
				break
			}
		}
		//create table
		if !create_flag {
			message.Println(string(list.name) + " createed.")
			continue
		}
		message.Println("create database table for", string(list.name))
		back, err := sql.Exec(list.cmd)
		if err != nil {
			return err
		}
		fmt.Println(back.LastInsertId())
		fmt.Println(back.RowsAffected())
		fmt.Println(list.cmd)

	}
	return nil
}

func tableList(sqltype string, db *sql.DB) ([]string, error) {

	switch sqltype {
	case "mysql":
	case "sqlite3":
		return sqlite3_table_list(db)
	default:

	}
	return []string{}, errors.New("Don't search sql Type")
}

func sqlite3_table_list(db *sql.DB) ([]string, error) {
	output := []string{}
	cmd := "select name from sqlite_master where type='table'"
	rows, err := db.Query(cmd)
	if err != nil {
		return []string{}, nil
	}
	for rows.Next() {
		str := ""
		err1 := rows.Scan(&str)
		if err1 != nil {
			return []string{}, err1
		}
		output = append(output, str)
	}
	return output, nil

}

func (t *createTableS) createTableCmd(sqltype string) error {
	output := "CREATE TABLE IF NOT EXISTS" + " "
	if t.name == "" {
		return errors.New("Don't input name data")
	}
	output += "\"" + string(t.name) + "\""
	output += " ("
	if t.st == nil {
		return errors.New("Don't input st data")
	}
	rt := reflect.TypeOf(t.st)
	count := 0
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		tmp := ""
		if i != 0 {
			output += ", "
		}
		switch f.Type.Kind() {
		case reflect.Int:
			tmp = f.Tag.Get("db")
			output += "\"" + tmp + "\" INTEGER"
		case reflect.String:
			tmp = f.Tag.Get("db")
			output += "\"" + tmp + "\" varchar"
		case reflect.Struct:
			tmp = f.Tag.Get("db")
			output += "\"" + tmp + "\" datetime"
		}
		if tmp == "id" {
			output += " PRIMARY KEY AUTOINCREMENT NOT NULL"
			count++
		} else if tmp == "created_at" {
			output += " NOT NULL"
			count++
		} else if tmp == "updated_at" {
			output += " NOT NULL"
			count++
		} else if tmp == "" {
			return errors.New("Don't tag setup for " + f.Name)
		}
	}
	if count != 3 {
		return errors.New("Don't Struct data for \"id\" and \"created_at\", \"updated_at\"")
	}
	output += ")"
	t.cmd = output
	return nil

}

func dballread(sqltype string, rows *sql.Rows, t_name Tablename) (interface{}, error) {

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
		ft := rt.Field(i)
		switch ft.Type.Kind() {
		case reflect.Int:
			i := int64(0)
			loaddata = append(loaddata, &i)
		case reflect.String:
			str := string("")
			loaddata = append(loaddata, &str)
		case reflect.Struct:
			// str := string("")
			var str time.Time
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
		case *time.Time:
			tmp := p.(*time.Time)
			output = append(output, *tmp)
		}
	}

	return output, err
}
