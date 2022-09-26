package table

import (
	"bookserver/message"
	"database/sql"
	"errors"
	"reflect"
	"strconv"
	"time"
)

const (
	Basefolder = "./db/"
	TimeLayout = "2006-01-02 15:04:05.999999999"
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
func (cfg *Config) sqlite3Add(t_name Tablename, v map[string]string) error {
	var err error
	var cmd string
	switch tablemap[t_name].(type) {
	case booknames:
		v["id"] = cfg.sqlite3IdMax(t_name)
		cmd, err = sqlite3RecodeAdd(t_name, booknames{}, v)
	case copyfile:
	case filelists:
	default:
		err = errors.New("not found db table")

	}
	if cmd != "" {
		message.Println(cmd)
		_, err = cfg.db.Exec(cmd)
	}
	return err
}

func (cfg *Config) sqlite3IdMax(t_name Tablename) string {
	id := 0
	cmd := "select max(id) from " + string(t_name)
	rows, err := cfg.db.Query(cmd)
	if err != nil {
		return ""
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&id)
	id++
	str := strconv.Itoa(id)
	if err != nil {
		return str
	}
	return str

}

func sqlite3RecodeAdd(t_name Tablename, tabletype interface{}, v map[string]string) (string, error) {
	cmd := "INSERT INTO " + string(t_name) + " "
	cmd_colume := ""
	cmd_vaule := ""
	now := time.Now()
	st := reflect.TypeOf(tabletype)
	for i := 0; i < st.NumField(); i++ {
		f := st.Field(i)
		if key := f.Tag.Get("db"); key != "" {
			var vaule string
			if i != 0 {
				cmd_colume += ","
				cmd_vaule += ","
			}
			cmd_colume += key
			vaule = v[key]
			switch f.Type.Kind() {
			case reflect.Int:
				cmd_vaule += vaule
			case reflect.String:
				cmd_vaule += "'" + vaule + "'"
			case reflect.Struct:
				cmd_vaule += "'" + now.Format(TimeLayout) + "'"
			}
		}
	}
	if cmd_colume == "" || cmd_vaule == "" {
		return "", errors.New("Don't created command")
	} else {
		cmd += " (" + cmd_colume + ") " + "VALUES" + " (" + cmd_vaule + ")"
	}
	return cmd, nil
}

//Update処理
func (cfg *Config) sqlUpdate(t_name Tablename, v map[string]string) error {
	return nil
}

//削除処理
func (cfg *Config) sqlite3Delete(t_name Tablename, v map[string]string) error {
	return nil
}

//読み取り処理v2
func (cfg *Config) sqlite3ReadV2(t_name Tablename, keyword map[string]string, keytype KeyWordOption) ([]map[string]string, error) {
	output := []map[string]string{}
	cmd := "SELECT * FROM" + " " + string(t_name)
	cmd += " " + "WHERE" + " " + convertCmdV2(t_name, keyword, keytype)
	rows, err := cfg.db.Query(cmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		data, err := sqlite3RowsRead(rows, t_name)
		if err != nil {
			return output, err
		}
		output = append(output, data)

	}

	return output, nil
}

func sqlite3RowsRead(rows *sql.Rows, t_name Tablename) (map[string]string, error) {
	var output map[string]string
	data, err := sqlite3RowsReadData(t_name)
	if err != nil {
		return output, err
	}
	err = rows.Scan(data...)
	if err != nil {
		return output, err
	}
	output, err = sqliteReadConvertV2(t_name, data...)

	return output, nil
}
func sqliteReadConvertV2(t_name Tablename, data ...interface{}) (map[string]string, error) {

	output := map[string]string{}
	table := tablemap[t_name]
	rt := reflect.TypeOf(table)
	if len(data) != rt.NumField() {
		return output, errors.New("Input data count err " + strconv.Itoa(len(v)))
	}
	for i := 0; i < rt.NumField(); i++ {
		if data[i] != nil {
			f := rt.Field(i)
			switch data[i].(type) {
			case *int64:
				tmp := data[i].(*int64)
				output[f.Name] = strconv.Itoa(int(*tmp))
			case *string:
				tmp := data[i].(*string)
				output[f.Name] = *tmp
			case *time.Time:
				ptmp := data[i].(*time.Time)
				tmp := *ptmp

				output[f.Name] = tmp.String()
			}
			// output[f.Name] =
		}
	}
	return output, nil
}

func sqlite3RowsReadData(t_name Tablename) ([]interface{}, error) {
	var output []interface{}
	loaddata := tablemap[t_name]
	if loaddata == nil {
		return output, errors.New("Don't input tablemap data for" + string(t_name))
	}
	rt := reflect.TypeOf(loaddata)
	for i := 0; i < rt.NumField(); i++ {
		ft := rt.Field((i))
		switch ft.Type.Kind() {
		case reflect.Int:
			i := int64(0)
			output = append(output, &i)
		case reflect.String:
			str := string("")
			output = append(output, &str)
		case reflect.Struct:
			// str := string("")
			var str time.Time
			output = append(output, &str)
		default:
			return output, errors.New("Error data type " + strconv.Itoa(int(ft.Type.Kind())))
		}
	}
	return output, nil
}

func convertCmdV2(t_name Tablename, keyword map[string]string, keytype KeyWordOption) string {
	output := ""
	if tablemap[t_name] == nil {
		return output
	}
	st := reflect.TypeOf(tablemap[t_name])
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
					output += "%" + keyword[f.Tag.Get("db")] + "%"
				case reflect.String:
					output += "'%" + keyword[f.Tag.Get("db")] + "%'"
				}
			}
			count++
		}
	}

	return output
}

//読み取り処理
func (cfg *Config) sqlite3_Read(t_name Tablename, keyword map[string]interface{}, keytype KeyWordOption) ([]any, error) {
	cmd := "SELECT * FROM" + " " + string(t_name)
	cmd += " " + "WHERE" + " " + convertCmd(keyword, keytype)
	var output []interface{}
	rows, err := cfg.db.Query(cmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
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
	defer rows.Close()
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
	var layout2 = TimeLayout
	if output, err = time.Parse(layout1, timdata); err != nil {
		if output, err = time.Parse(layout2, timdata); err != nil {
			return output, err
		}
	}
	return output, nil
}
