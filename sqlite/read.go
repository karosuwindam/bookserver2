package sqlite

import (
	"errors"
	"reflect"
)

type KeyWordOption string //検索オプション

//検索オプションの値
//
//AND keyword=data and
//OR keyword=data or
//AND_Like keyword like %keyword% and
//OR_LIKE keyword like %keyword% or
const (
	AND      KeyWordOption = "and"
	OR       KeyWordOption = "or"
	AND_Like KeyWordOption = "and_like"
	OR_Like  KeyWordOption = "or_like"
)

func (t *sqliteConfig) Read(tname string, stu interface{}, slice interface{}, v map[string]string, keytype KeyWordOption) error {
	cmd, err := createReadCmd(tname, stu, v, keytype)
	if err != nil {
		return err
	}
	rows, err := t.db.Query(cmd)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {

	}

	return err
}

func createReadCmd(tname string, stu interface{}, keyword map[string]string, keytype KeyWordOption) (string, error) {
	rt := reflect.TypeOf(stu)
	if rt.Kind() == reflect.UnsafePointer {
		return "", errors.New("This input stu data is pointer")
	}
	cmd := "SELECT * FROM" + " " + tname
	if len(keyword) == 0 {
		return cmd, nil
	}
	cmd += " " + "WHERE" + " " + convertCmd(stu, keyword, keytype)

	return cmd, nil

}

func convertCmd(stu interface{}, keyword map[string]string, keytype KeyWordOption) string {
	output := ""
	if stu == nil {
		return output
	}
	st := reflect.TypeOf(stu)
	count := 0
	for i := 0; i < st.NumField(); i++ {
		f := st.Field(i)
		if keyword[f.Tag.Get("db")] != "" {
			if count != 0 {
				switch keytype {
				case AND_Like:
					output += " " + string(AND) + " "
				case OR_Like:
					output += " " + string(OR) + " "
				default:
					output += " " + string(keytype) + " "
				}

			}
			if keytype == AND || keytype == OR {
				output += f.Tag.Get("db") + "="
				switch f.Type.Kind() {
				case reflect.Int:
					output += keyword[f.Tag.Get("db")]
				case reflect.String:
					output += "'" + keyword[f.Tag.Get("db")] + "'"
				}
			} else {
				output += f.Tag.Get("db") + " like "
				switch f.Type.Kind() {
				case reflect.Int:
					output += "'%" + keyword[f.Tag.Get("db")] + "%'"
				case reflect.String:
					output += "'%" + keyword[f.Tag.Get("db")] + "%'"
				}
			}
			count++
		}
	}

	return output
}
