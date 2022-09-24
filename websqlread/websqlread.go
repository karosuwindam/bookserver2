package websqlread

import (
	"bookserver/message"
	"bookserver/table"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type WebSqlRead struct {
	sql       *sql.DB
	sqlconfig *table.Config
	msg       message.Message
	table     []table.Tablename
}

//URL解析
func urlAnalysis(url string) []string {
	tmp := []string{}
	for _, str := range strings.Split(url[1:], "/") {
		tmp = append(tmp, str)
	}
	return tmp
}

//セットアップ
func Setup(sql *table.Config) (*WebSqlRead, error) {
	var err error = nil
	output := &WebSqlRead{}
	output.sqlconfig = sql
	output.table = []table.Tablename{table.Booknames, table.Copyfile, table.Filelists}
	output.msg = message.Message{Name: "sqlread", Status: "OK", Code: http.StatusOK}
	return output, err
}

//メッセージのバック
func (t *WebSqlRead) outputmessage(w http.ResponseWriter) {
	fmt.Fprintf(w, "%v", t.msg.Output())
}

//Getread処理
func (t *WebSqlRead) getSqlRead(id int, table table.Tablename, w http.ResponseWriter, r *http.Request) {
	result := message.Result{Name: "SQL Read", Date: time.Now()}
	if table == "" {
		t.outputmessage(w)
		return
	} else {
		if id == -1 {
			result.Option = ""
			result.Result = "input id data error"

		} else {
			result.Option = "id=" + strconv.Itoa(id)
		}
		fmt.Fprintf(w, "%v", result.Output())
		// fmt.Fprintf(w, "%v", result.Result)

	}

}

//Getread All処理
func (t *WebSqlRead) getSqlReadAll(table table.Tablename, w http.ResponseWriter, r *http.Request) {
	result := message.Result{Name: "SQL Read", Date: time.Now()}
	if table == "" {
		t.outputmessage(w)
		return
	} else {
		result.Option = "all"
		data, err := t.sqlconfig.ReadAll(table)
		if err != nil {
			result.Result = err.Error()
		} else {
			result.Result = data
		}

		fmt.Fprintf(w, "%v", result.Output())

	}

}

//基本処置
func (t *WebSqlRead) sqlread_defult(w http.ResponseWriter, r *http.Request) {
	result := message.Result{Name: "SQL Read", Date: time.Now(), Result: "OK"}
	fmt.Fprintf(w, "%v", result.Output())
}

//Method別処理
func Websqlread(t *WebSqlRead, w http.ResponseWriter, r *http.Request) {
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
