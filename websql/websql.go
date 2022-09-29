package websql

import (
	"bookserver/message"
	"bookserver/table"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

type WebSql struct {
	sql       *sql.DB
	sqlconfig *table.Config
	rst       message.Result
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
func Setup(sql *table.Config) (*WebSql, error) {
	var err error = nil
	output := &WebSql{}
	output.sqlconfig = sql
	output.table = []table.Tablename{table.Booknames, table.Copyfile, table.Filelists}
	output.rst = message.Result{Name: "SQL Read", Code: http.StatusOK}
	return output, err
}

//メッセージのバック
func (t *WebSql) outputmessage(w http.ResponseWriter) {
	w.WriteHeader(t.rst.Code)
	fmt.Fprintf(w, "%v", t.rst.Output())
}
