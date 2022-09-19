package table

import (
	"bookserver/config"
	"log"
)

type Config struct {
}

//基本の設定
func Setup(data *config.Config) (*Config, error) {
	return &Config{}, nil
}

//DBを開く
func (cfg *Config) Open() {
	log.Println("SQL server open")
}

//DBを閉じる
func (cfg *Config) Close() {
	log.Println("SQL server close")

}
