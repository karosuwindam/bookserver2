package sqlite

import (
	"errors"
	"reflect"
)

func (t *sqliteConfig) CreateTable(tname string, stu interface{}) error {
	cmd, err := createTableCmd(tname, stu)
	if err != nil {
		return err
	}
	_, err = t.db.Exec(cmd)

	return err
}

func (t *sqliteConfig) ReadTableList() ([]string, error) {
	var output []string
	cmd, err := readTableAllCmd()
	if err != nil {
		return output, err
	}
	rows, err := t.db.Query(cmd)
	if err != nil {
		return output, err
	}
	defer rows.Close()
	for rows.Next() {
		str := ""
		err = rows.Scan(&str)
		if err != nil {
			return []string{}, err
		}
		output = append(output, str)
	}

	return output, err
}

func (t *sqliteConfig) ReadCreateTableCmd(tname string) (string, error) {
	var output string
	cmd, err := readCreateTableCmd(tname)
	if err != nil {
		return output, err
	}
	rows, err := t.db.Query(cmd)
	if err != nil {
		return output, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&output)
		if err != nil {
			return "", err
		}
	}

	return output, err

}

func (t *sqliteConfig) DropTable(tname string) error {
	cmd, err := dropTableCmd(tname)
	if err != nil {
		return err
	}
	_, err = t.db.Exec(cmd)
	return err

}

func createTableCmd(tname string, stu interface{}) (string, error) {
	cmd := "CREATE TABLE IF NOT EXISTS" + " "
	if tname == "" {
		return "", errors.New("Don't input name data")
	}
	cmd += "\"" + tname + "\""
	cmd += " ("
	if reflect.TypeOf(stu).Kind() != reflect.Struct {
		return "", errors.New("Don't input st data")
	}
	rt := reflect.TypeOf(stu)
	count := 0
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		tmp := ""
		if i != 0 {
			cmd += ", "
		}
		switch f.Type.Kind() {
		case reflect.Int:
			tmp = f.Tag.Get("db")
			cmd += "\"" + tmp + "\" INTEGER"
		case reflect.String:
			tmp = f.Tag.Get("db")
			cmd += "\"" + tmp + "\" varchar"
		}
		if tmp == "id" {
			cmd += " PRIMARY KEY AUTOINCREMENT NOT NULL"
			count++
		} else if tmp == "" {
			return "", errors.New("Don't tag setup for " + f.Name)
		}
	}
	if count == 0 {
		return "", errors.New("Don't Struct data for \"id\"")
	}
	cmd += ", \"created_at\" datetime"
	cmd += ", \"updated_at\" datetime"
	cmd += ")"
	return cmd, nil

}

func dropTableCmd(tname string) (string, error) {
	cmd := "DROP TABLE IF EXISTS" + " '" + tname + "'"
	return cmd, nil

}

func readTableAllCmd() (string, error) {
	cmd := "SELECT name FROM sqlite_master WHERE type='table'"
	return cmd, nil

}

func readCreateTableCmd(tname string) (string, error) {
	cmd := "SELECT sql FROM sqlite_master WHERE type='table' AND name='" + tname + "'"
	return cmd, nil
}
