package table

import (
	"bookserver/config"
	"database/sql"
	"errors"
	"log"
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
	output.Db_file = "development.sqlite3"

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
		log.Println(msg)
		cfg.Message = msg

		// data, _ := cfg.ReadAll(Copyfile)
		// fmt.Println(data)
		// if jsondata, err1 := json.Marshal(data); err1 == nil {
		// 	fmt.Println(string(jsondata))

		// }
		// cfg.sqlite3_close()
	}
	return err
}

//DBを閉じる
func (cfg *Config) Close() {
	msg := "SQL server close"
	log.Println(msg)
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

//テーブル内のカラムの検索による読み取り
//
// v map[]interface{} = ["カラムのkeyword"]{検索の値}
func (cfg *Config) Search(t_name Tablename, v map[string]interface{}) ([]any, error) {
	switch cfg.Db_name {
	case "mysql":
	case "sqlite3":
		return cfg.sqlite3_Read(t_name, v, OR_Like)
	default:

	}
	return nil, errors.New("Don't select db type")
}
