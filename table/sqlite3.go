package table

import (
	"database/sql"
)

func (cfg *Config) sqlite3_open() error {
	var err error
	cfg.db, err = sql.Open("sqlite3", cfg.Db_file)
	return err

}

func (cfg *Config) sqlite3_close() {
	cfg.db.Close()
}

// func (cfg *Config) sqlite3_ReadAll(v ...any) error {
// 	//メタデータ解析
// 	var readdata []interface{}
// 	rt := reflect.TypeOf(v[0])
// 	for i := 0; i < rt.NumField(); i++ {
// 		f := rt.Field(i)
// 		readdata = append(readdata, checkdata(f))
// 		fmt.Println(f.Name, f.Type, f.Tag.Get("db"), f.Tag.Get("type"))
// 	}
// 	//sql構文作成
// 	cmd := "SELECT * FROM " + "booknames"
// 	//sqlによる読み込み
// 	rows, err := cfg.db.Query(cmd)
// 	if err != nil {
// 		return err
// 	}
// 	for rows.Next() {
// 		rows.Scan(v...)
// 	}
// 	return nil
// }

// func checkdata(v reflect.StructField) any {
// 	switch v.Type.Kind() {
// 	case reflect.Int:
// 		return int64(0)
// 	case reflect.String:
// 		return string("")
// 	case reflect.Struct:
// 		return time.Time{}
// 	default:
// 		fmt.Println(v.Type.Kind())
// 		return int(0)
// 	}
// }

func (cfg *Config) sqlite3_ReadAll(t_name string) ([]any, error) {
	cmd := "SELECT * FROM " + t_name
	var output []interface{}
	rows, err := cfg.db.Query(cmd)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tmp, err := dbread(cfg.Db_name, rows, t_name)
		if err != nil {
			return output, err
		}
		output = append(output, tmp)
	}

	return output, nil
}
