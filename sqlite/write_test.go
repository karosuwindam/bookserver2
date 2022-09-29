package sqlite

import (
	"os"
	"testing"
)

func TestTableWrite(t *testing.T) {

	type TableTest struct {
		id  int    `db:"id"`
		str string `db:"str"`
		i   int    `db:"i"`
	}

	testtablename := "test"
	testdbname := "test.db"

	sql := Setup(testdbname)
	_ = sql.Open()
	defer sql.Close()
	sql.CreateTable(testtablename, TableTest{})
	defer os.Remove(testdbname)

	wdata := TableTest{id: 10, str: "data", i: 500}
	t.Log(wdata)
	err := sql.Add(testtablename, &wdata)
	if err != nil {
		t.Errorf("Don't Added %v Table", testtablename)
	}

	i := sql.sqlite3IdMax(testtablename)
	if i != 2 {
		t.Errorf("%v table count+1 = %v", testtablename, i)
	}
}
