package websql

import (
	"bookserver/table"
	"fmt"
	"log"
	"net/http"
	"time"
)

//Getread idによる読み取り
func (t *WebSql) getSqlSerch(keyword string, table table.Tablename, w http.ResponseWriter, r *http.Request) {
	t.rst.Date = time.Now()
	if table == "" {
		t.rst.Code = http.StatusNoContent
		t.outputmessage(w)
		return
	} else {
		t.rst.Option = "keyword=" + keyword
		skeyword := map[string]interface{}{
			"keyword": keyword,
		}
		if keyword != "" {
			data, err := t.sqlconfig.Search(table, skeyword)
			if err != nil {
				log.Println(err.Error())
				t.rst.Result = err.Error()
			} else {
				t.rst.Result = data
			}
		}

		// fmt.Fprintf(w, "%v", result.Output())
		t.outputmessage(w)

	}

}

//Method別処理
func webSqlSerch(t *WebSql, w http.ResponseWriter, r *http.Request) {
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
		t.rst.Code = http.StatusNoContent
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
	t.rst.Code = http.StatusOK
	t.rst.Date = time.Now()
	t.getSqlSerch(urldata[point_url+2], table, w, r)

}

func WebsqlSerch(it interface{}, w http.ResponseWriter, r *http.Request) {
	switch it.(type) {
	case *WebSql:
		t := it.(*WebSql)
		webSqlSerch(t, w, r)

	default:
		log.Println("input point type err")
		w.WriteHeader(400)
		fmt.Fprintf(w, "input point type err")
	}

}
