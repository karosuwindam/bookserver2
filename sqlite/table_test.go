package sqlite

import (
	"testing"
)

type TableTest struct {
	id  int    `db:"id"`
	str string `db:"str"`
	i   int    `db:"i"`
}

func createchcmdTableTest(name string) string {
	cmd := "CREATE TABLE IF NOT EXISTS \"" + name + "\" "
	cmd += "(" + "\"id\" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL"
	cmd += ", \"str\" varchar"
	cmd += ", \"i\" INTEGER"
	cmd += ", \"created_at\" datetime, \"updated_at\" datetime)"
	return cmd
}

func TestCreateTable(t *testing.T) {
	sql := Setup("test.db")
	err := sql.Open()
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	err = sql.Close()
	if err != nil {
		t.Errorf("%v", err.Error())

	}
	str, err1 := createTableCmd("test", TableTest{})
	if err1 != nil {
		t.Errorf("%v", err1.Error())
	}
	if str != createchcmdTableTest("test") {
		t.Errorf("output:%v\ncheck:%v", str, createchcmdTableTest("test"))
	}

	t.Logf("%v", str)

}
