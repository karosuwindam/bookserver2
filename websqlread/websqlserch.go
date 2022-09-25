package websqlread

import (
	"bookserver/message"
	"bookserver/table"
	"fmt"
	"log"
	"net/http"
	"time"
)

//Getread idによる読み取り
func (t *WebSqlRead) getSqlSerch(keyword string, table table.Tablename, w http.ResponseWriter, r *http.Request) {
	result := message.Result{Name: "SQL Read", Date: time.Now()}
	if table == "" {
		t.outputmessage(w)
		return
	} else {
		result.Option = "keyword=" + keyword
		skeyword := map[string]interface{}{
			"keyword": keyword,
		}
		if keyword != "" {
			data, err := t.sqlconfig.Search(table, skeyword)
			if err != nil {
				result.Result = err.Error()
			} else {
				result.Result = data
			}
		}

		fmt.Fprintf(w, "%v", result.Output())

	}

}

//Method別処理
func webSqlSerch(t *WebSqlRead, w http.ResponseWriter, r *http.Request) {
	urldata := urlAnalysis(r.URL.Path)
	point_url := 0
	var table table.Tablename
	for i, str := range urldata {
		if str == "serch" {
			point_url = i
			break
		}
	}
	if point_url == 0 || len(urldata[point_url:]) < 3 {
		t.sqlread_defult(w, r)
		return
	} else {
		for _, str := range t.table {
			if string(str) == urldata[point_url+1] {
				table = str
				break
			}
		}
	}
	t.getSqlSerch(urldata[point_url+2], table, w, r)

}

func WebsqlSerch(it interface{}, w http.ResponseWriter, r *http.Request) {
	switch it.(type) {
	case *WebSqlRead:
		t := it.(*WebSqlRead)
		webSqlSerch(t, w, r)

	default:
		log.Println("input point type err")
		w.WriteHeader(400)
		fmt.Fprintf(w, "input point type err")
	}

}
