package websql

import (
	"bookserver/message"
	"bookserver/table"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

//Getread idによる読み取り
func (t *WebSql) getSqlRead(id int, table table.Tablename, w http.ResponseWriter, r *http.Request) {
	t.rst.Result = nil
	if table == "" {
		t.rst.Code = http.StatusNoContent
		t.outputmessage(w)
		return
	} else {
		if id == -1 {
			t.rst.Code = http.StatusNoContent
			t.rst.Option = ""
			t.rst.Result = "input id data error"

		} else {
			t.rst.Option = "id=" + strconv.Itoa(id)
			keyword := map[string]interface{}{
				"id": id,
			}
			data, err := t.sqlconfig.Read(table, keyword)
			if err != nil {
				log.Println(err.Error())
				t.rst.Result = err.Error()
			} else {
				t.rst.Result = data
			}
		}
		t.outputmessage(w)
		// fmt.Fprintf(w, "%v", result.Output())

	}

}

//Getread All処理
func (t *WebSql) getSqlReadAll(table table.Tablename, w http.ResponseWriter, r *http.Request) {
	t.rst = message.Result{Name: "SQL Read", Date: time.Now(), Code: t.rst.Code, Option: t.rst.Option, Result: t.rst.Result}
	if table == "" {
		t.rst.Code = http.StatusNoContent
		t.outputmessage(w)
		return
	} else {
		t.rst.Option = "all"
		data, err := t.sqlconfig.ReadAll(table)
		if err != nil {
			log.Println(err.Error())
			t.rst.Result = err.Error()
		} else {
			t.rst.Result = data
		}

		fmt.Fprintf(w, "%v", t.rst.Output())

	}

}

//基本処置
func (t *WebSql) sqlread_defult(w http.ResponseWriter, r *http.Request) {
	t.rst.Date = time.Now()
	t.rst.Result = nil
	t.outputmessage(w)
}

//Method別処理
func websqlread(t *WebSql, w http.ResponseWriter, r *http.Request) {
	urldata := urlAnalysis(r.URL.Path)
	point_url := 0
	var table table.Tablename
	id := -1
	for i, str := range urldata {
		if str == "read" {
			point_url = i
			break
		}
	}
	if point_url == 0 {
		t.rst.Code = http.StatusBadRequest
		t.sqlread_defult(w, r)
		return
	} else {
		for _, str := range t.table {
			if string(str) == urldata[point_url+1] {
				table = str
				break
			}
		}
		if len(urldata[point_url:]) < 3 {
			id = -1
		} else {
			i, err := strconv.Atoi(urldata[point_url+2])
			if err != nil {
				id = -1
			} else {
				id = i
			}
		}
	}
	t.rst.Code = http.StatusOK
	t.rst.Date = time.Now()
	t.rst.Name = "SQL Read"
	t.rst.Option = ""
	switch r.Method {
	case "GET":
		t.getSqlRead(id, table, w, r)
	case "POST":
		t.sqlread_defult(w, r)
	case "LIST":
		t.getSqlReadAll(table, w, r)
	default:
		t.sqlread_defult(w, r)
	}

}

func Websqlread(it interface{}, w http.ResponseWriter, r *http.Request) {
	switch it.(type) {
	case *WebSql:
		t := it.(*WebSql)
		websqlread(t, w, r)

	default:
		log.Println("input point type err")
		w.WriteHeader(400)
		fmt.Fprintf(w, "input point type err")
	}

}
