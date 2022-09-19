package table

import "time"

type booknames struct {
	Id         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Title      string    `json:"title" db:"title"`
	Writer     string    `json:"writer db:"writer"`
	Brand      string    `json:"brand" db:"brand"`
	Booktype   string    `json:"booktype" db:"booktype"`
	Ext        string    `json:"ext" db:"ext"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
}

func Read() {

}
