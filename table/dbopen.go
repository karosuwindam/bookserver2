package table

import (
	"bookserver/config"
	"bookserver/message"
	"database/sql"
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

var tablemap map[Tablename]interface{}

type Config struct {
	Db_name string  `json:"name"`     //sqlの種類
	Db_host string  `json:"host"`     //sqlの接続IPやDNS
	Db_port string  `json:"port"`     //SQLの接続port
	Db_user string  `json:"user"`     //SQLの接続USER
	Db_pass string  `json:"pasword"`  //SQLの接続パス
	Db_file string  `json:"filepass"` //SQLデータベース接続ファイルパス
	db      *sql.DB //開いたSQLについて
	Message string  //Helth checkに渡す用
}

type SerchID struct {
	Id int `db:"id"`
}

//基本の設定
func Setup(data *config.Config) (*Config, error) {
	output := &Config{
		Db_name: data.Sql.DBNAME,
		Db_host: data.Sql.DBHOST,
		Db_port: data.Sql.DBPORT,
		Db_user: data.Sql.DBUSER,
		Db_pass: data.Sql.DBPASS,
		Db_file: data.Sql.DBFILE,
	}

	//Defult
	output.Db_name = "sqlite3"
	// output.Db_file = "development.sqlite3"
	// output.Db_file = "test3.db"

	//テーブル名に対するテーブルの型
	tablemap = map[Tablename]interface{}{
		Booknames: booknames{},
		Copyfile:  copyfile{},
		Filelists: filelists{},
	}

	return output, nil
}

//DBを開く
func (cfg *Config) Open() error {
	err := error(nil)
	switch cfg.Db_name {
	case "mysql":
	case "sqlite3":
		err = cfg.sqlite3_open()
	default:
	}
	if err == nil {
		msg := "SQL server open"
		message.Println(msg)
		cfg.Message = msg

	}
	return err
}

//DBを閉じる
func (cfg *Config) Close() {
	msg := "SQL server close"
	message.Println(msg)
	cfg.Message = msg
}

//Table作成
func (cfg *Config) Create_Table() error {
	switch cfg.Db_name {
	case "mysql":
	case "sqlite3":
		return cfg.sqlite3_table_create()
	default:

	}

	return errors.New("Don't create db table")
}

//Tableリストの読み取り
func (cfg *Config) List_Table() ([]string, error) {
	return tableList(cfg.Db_name, cfg.db)
}

//テーブル内のすべてのデータ読み取り
func (cfg *Config) ReadAll(t_name Tablename) ([]any, error) {
	switch cfg.Db_name {
	case "mysql":
	case "sqlite3":
		return cfg.sqlite3_ReadAll(t_name)
	default:

	}
	return nil, errors.New("Don't select db type")
}

//テーブル内のレコードを追加
//
//v map[string]interface{} = [設定名]{登録の値}
func (cfg *Config) Add(t_name Tablename, v map[string]string) error {
	switch cfg.Db_name {
	case "mysql":
	case "sqlite3":
		return cfg.sqlite3Add(t_name, v)
	default:

	}
	return errors.New("Don't select db type")
}

//テーブル内の特定IDのレコードの更新
//
//v map[string]interface{} = [設定名]{登録の値}
func (cfg *Config) Update(t_name Tablename, v map[string]string) error {
	switch cfg.Db_name {
	case "mysql":
	case "sqlite3":
		return cfg.sqlUpdate(t_name, v)
	default:

	}
	return errors.New("Don't select db type")

}

//テーブル内の特定IDのレコードを削除
//
//v map[string]interface{} = [設定名]{登録の値}
func (cfg *Config) Delete(t_name Tablename, v map[string]string) error {
	switch cfg.Db_name {
	case "mysql":
	case "sqlite3":
		return cfg.sqlite3Delete(t_name, v)
	default:

	}
	return errors.New("Don't select db type")

}

//テーブル内の特定カラムによる読み取り.
//
// v map[]interface{} = ["カラムのkeyword"]{検索の値}
func (cfg *Config) Read(t_name Tablename, v map[string]interface{}) ([]any, error) {
	switch cfg.Db_name {
	case "mysql":
	case "sqlite3":
		return cfg.sqlite3_Read(t_name, v, AND)
	default:

	}
	return nil, errors.New("Don't select db type")
}

func (cfg *Config) Readv2(t_name Tablename, v map[string]string) ([]map[string]string, error) {
	switch cfg.Db_name {
	case "mysql":
	case "sqlite3":
		return cfg.sqlite3ReadV2(t_name, v, AND)
	default:

	}
	return nil, errors.New("Don't select db type")
}

//テーブル内のカラムの検索による読み取り
//
// v map[]interface{} = ["カラムのkeyword"]{検索の値}
func (cfg *Config) Search(t_name Tablename, v map[string]interface{}) ([]any, error) {
	switch cfg.Db_name {
	case "mysql":
	case "sqlite3":
		keyword := createSerchKeyword(t_name, v)
		return cfg.sqlite3_Read(t_name, keyword, OR_Like)
	default:

	}
	return nil, errors.New("Don't select db type")
}
func (cfg *Config) SearchV2(t_name Tablename, v map[string]string) ([]map[string]string, error) {
	switch cfg.Db_name {
	case "mysql":
	case "sqlite3":
		keyword := createSerchKeywordV2(t_name, v)
		return cfg.sqlite3ReadV2(t_name, keyword, OR_Like)
	default:

	}
	return nil, errors.New("Don't select db type")
}

func createSerchKeyword(t_name Tablename, keyword map[string]interface{}) map[string]interface{} {
	output := map[string]interface{}{}
	ts := tablemap[t_name]
	st := reflect.TypeOf(ts)
	for keyname, data := range keyword {
		if keyname == "keyword" {
			for i := 0; i < st.NumField(); i++ {
				f := st.Field(i)
				switch f.Type.Kind() {
				case reflect.String:
					tagname := f.Tag.Get("db")
					if tagname != "" {
						output[tagname] = data
					}
				}
			}
		} else {
			output[keyname] = data
		}
	}
	return output
}
func createSerchKeywordV2(t_name Tablename, keyword map[string]string) map[string]string {
	output := map[string]string{}
	ts := tablemap[t_name]
	st := reflect.TypeOf(ts)
	for keyname, data := range keyword {
		if keyname == "keyword" {
			for i := 0; i < st.NumField(); i++ {
				f := st.Field(i)
				switch f.Type.Kind() {
				case reflect.String:
					tagname := f.Tag.Get("db")
					if tagname != "" {
						output[tagname] = data
					}
				}
			}
		} else {
			output[keyname] = data
		}
	}
	return output
}
