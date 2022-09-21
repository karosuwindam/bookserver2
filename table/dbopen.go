package table

import (
	"bookserver/config"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type Config struct {
	Db_name string  `json:"name"`     //sqlの種類
	Db_host string  `json:"host"`     //sqlの接続IPやDNS
	Db_port string  `json:"port"`     //SQLの接続port
	Db_user string  `json:"user"`     //SQLの接続USER
	Db_pass string  `json:"pasword"`  //SQLの接続パス
	Db_file string  `json:"filepass"` //SQLデータベース接続ファイルパス
	db      *sql.DB //開いたSQLについて
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
	output.Db_file = "./db/development.sqlite3"
	// output.Db_file = "./db/test1.db"

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
		log.Println("SQL server open")

		data, _ := cfg.ReadAll(string(Booknames))
		fmt.Println(data)
		if jsondata, err1 := json.Marshal(data); err1 == nil {
			fmt.Println(string(jsondata))

		}
		cfg.sqlite3_close()
	}
	return err
}

//DBを閉じる
func (cfg *Config) Close() {
	log.Println("SQL server close")

}

func (cfg *Config) ReadAll(t_name string) ([]any, error) {
	switch cfg.Db_name {
	case "mysql":
	case "sqlite3":
		return cfg.sqlite3_ReadAll(t_name)
	default:

	}
	return nil, errors.New("Don't select db type")
}
