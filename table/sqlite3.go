package table

import (
	"database/sql"
	"strconv"
	"time"
)

const (
	Basefolder = "./db/"
)

func (cfg *Config) sqlite3_open() error {
	var err error
	cfg.db, err = sql.Open("sqlite3", Basefolder+cfg.Db_file)
	return err

}

func (cfg *Config) sqlite3_close() {
	cfg.db.Close()
}

func (cfg *Config) sqlite3_table_create() error {
	return tablecreate(cfg.Db_name, cfg.db)
}

//追加処理
func (cfg *Config) sqlite3_Add(t_name Tablename, v interface{}) error {
	return nil
}

func (cfg *Config) sqlite3_Read(t_name Tablename, keyword map[string]interface{}, keytype KeyWordOption) ([]any, error) {
	cmd := "SELECT * FROM" + " " + string(t_name)
	cmd += " " + "WHERE" + " " + convertCmd(keyword, keytype)
	var output []interface{}
	rows, err := cfg.db.Query(cmd)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tmp, err := dballread(cfg.Db_name, rows, t_name)
		if err != nil {
			return output, err
		}
		output = append(output, tmp)
	}

	return output, nil
}

func convertCmd(keyword map[string]interface{}, keytype KeyWordOption) string {
	output := ""
	count := 0
	for keyname, data := range keyword {
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
			output += keyname + "="
			switch data.(type) {
			case int:
				output += strconv.Itoa(data.(int))
			case string:
				output += "'" + data.(string) + "'"
			}
		} else {
			output += keyname + " like "
			switch data.(type) {
			case int:
				output += "%" + strconv.Itoa(data.(int)) + "%"
			case string:
				output += "'%" + data.(string) + "%'"
			}
		}
		count++
	}
	return output
}

func (cfg *Config) sqlite3_ReadAll(t_name Tablename) ([]any, error) {
	cmd := "SELECT * FROM " + string(t_name)
	var output []interface{}
	rows, err := cfg.db.Query(cmd)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tmp, err := dballread(cfg.Db_name, rows, t_name)
		if err != nil {
			return output, err
		}
		output = append(output, tmp)
	}

	return output, nil
}

//sqliteの時間文字列を変換
func timeconvert(timdata string) (time.Time, error) {
	var err error
	var output time.Time
	var layout1 = "2006-01-02T15:04:05.999999Z"
	var layout2 = "2006-01-02 15:04:05.999999999"
	if output, err = time.Parse(layout1, timdata); err != nil {
		if output, err = time.Parse(layout2, timdata); err != nil {
			return output, err
		}
	}
	return output, nil
}
