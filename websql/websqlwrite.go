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

//データ追加処理
func (t *WebSql) websqladd(table table.Tablename, w http.ResponseWriter, r *http.Request) {

	t.rst = message.Result{Name: "SQL Add", Date: time.Now(), Code: t.rst.Code, Result: t.rst.Result}
	if table == "" {
		t.rst.Code = http.StatusNoContent
		t.outputmessage(w)
		return
	} else {
		v := readFormData(r)
		if len(v) == 0 {
			t.rst.Result = "no data"
			fmt.Fprintf(w, "%v", t.rst.Output())

			return
		} else {
			count := 0
			for key, value := range v {
				if count != 0 {
					t.rst.Option += " "
				}
				t.rst.Option += key + "=" + value
				count++

			}
		}
		err := t.sqlconfig.Add(table, v)
		if err != nil {
			log.Println(err.Error())
			t.rst.Result = err.Error()
		} else {
			t.rst.Result = "OK"
		}

		fmt.Fprintf(w, "%v", t.rst.Output())

	}
}

func readFormData(r *http.Request) map[string]string {
	output := map[string]string{}

	err := r.ParseForm()
	if err != nil {
	}
	for k, v := range r.PostForm {
		output[k] = v[0]
	}
	return output
}

//データのアップデート処理
func (t *WebSql) websqlupdate(id int, table table.Tablename, w http.ResponseWriter, r *http.Request) {

	t.rst = message.Result{Name: "SQL Update", Date: time.Now(), Code: t.rst.Code, Option: t.rst.Option, Result: t.rst.Result}
	if table == "" {
		t.rst.Code = http.StatusNoContent
		t.outputmessage(w)
		return
	} else {

		fmt.Fprintf(w, "%v", t.rst.Output())

	}
}

//Method別処理
func websqlAdd(t *WebSql, w http.ResponseWriter, r *http.Request) {
	urldata := urlAnalysis(r.URL.Path)
	point_url := 0
	var table table.Tablename
	for i, str := range urldata {
		if str == "add" {
			point_url = i
			break
		}
	}
	if point_url == 0 || len(urldata[point_url:]) < 2 {
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
	}
	t.rst.Code = http.StatusOK
	t.rst.Date = time.Now()
	switch r.Method {
	case "POST":
		t.websqladd(table, w, r)
	default:
		t.sqlread_defult(w, r)
	}

}

func WebsqlAdd(it interface{}, w http.ResponseWriter, r *http.Request) {
	switch it.(type) {
	case *WebSql:
		t := it.(*WebSql)
		websqlAdd(t, w, r)

	default:
		log.Println("input point type err")
		w.WriteHeader(400)
		fmt.Fprintf(w, "input point type err")
	}

}

//Method別処理
func websqlUpdate(t *WebSql, w http.ResponseWriter, r *http.Request) {
	urldata := urlAnalysis(r.URL.Path)
	point_url := 0
	var table table.Tablename
	id := -1
	for i, str := range urldata {
		if str == "update" {
			point_url = i
			break
		}
	}
	if point_url == 0 || len(urldata[point_url:]) < 3 {
		t.rst.Code = http.StatusBadRequest
		t.sqlread_defult(w, r)
		return
	} else {
		for _, str := range t.table {
			if string(str) == urldata[point_url+1] {
				table = str
				break
			}
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
	switch r.Method {
	case "POST":
		t.websqlupdate(id, table, w, r)
	default:
		t.sqlread_defult(w, r)
	}

}

func WebsqlUpdate(it interface{}, w http.ResponseWriter, r *http.Request) {
	switch it.(type) {
	case *WebSql:
		t := it.(*WebSql)
		websqlAdd(t, w, r)

	default:
		log.Println("input point type err")
		w.WriteHeader(400)
		fmt.Fprintf(w, "input point type err")
	}

}
