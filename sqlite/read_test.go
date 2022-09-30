package sqlite

import (
	"os"
	"testing"
)

func TestTableRead(t *testing.T) {

	type TableTest struct {
		Id  int    `db:"id"`
		Str string `db:"str"`
		I   int    `db:"i"`
	}

	testtablename := "test"
	testdbname := "test.db"

	sql := Setup(testdbname)
	_ = sql.Open()
	defer sql.Close()
	sql.CreateTable(testtablename, TableTest{})
	defer os.Remove(testdbname)

	wdata := TableTest{Id: 1, Str: "data", I: 500}
	sql.Add(testtablename, &wdata)
	wdata1 := TableTest{Id: 1, Str: "data1", I: 45500}
	sql.Add(testtablename, &wdata1)
	wdata2 := TableTest{Id: 1, Str: "data2", I: 5500}
	sql.Add(testtablename, &wdata2)

	t.Log("-----------Read data ---------------")
	rdata := []TableTest{}
	err := sql.Read(testtablename, TableTest{}, &rdata, map[string]string{}, AND)
	if err != nil {
		t.Errorf("read err :%v", err.Error())
		t.FailNow()
	}
	ckdata := []TableTest{wdata, wdata1, wdata2}
	if len(ckdata) != len(rdata) {
		t.Errorf("read count ng %v=%v", len(ckdata), len(rdata))
		t.FailNow()
	}
	for i, tmp := range rdata {
		tmp2 := ckdata[i]
		if tmp.I != tmp2.I {
			t.Errorf("read ng for i %v=%v", tmp.I, tmp2.I)
		}
		if tmp.Str != tmp2.Str {
			t.Errorf("read ng for str %v=%v", tmp.Str, tmp2.Str)
		}
	}
	t.Log("-----------Read data Serch by ID---------------")
	rdata1 := []TableTest{}
	err1 := sql.Read(testtablename, TableTest{}, &rdata1, map[string]string{"id": "1"}, AND)
	if err1 != nil {
		t.Errorf("read err :%v", err1.Error())
	}
	if len(rdata) != 1 {
		t.Errorf("read count ng %v=%v", 1, len(rdata))
		t.FailNow()
	}
	for i, tmp := range rdata {
		tmp2 := ckdata[i]
		if tmp.I != tmp2.I {
			t.Errorf("read ng for i %v=%v", tmp.I, tmp2.I)
		}
		if tmp.Str != tmp2.Str {
			t.Errorf("read ng for str %v=%v", tmp.Str, tmp2.Str)
		}
	}
	t.Log("table data read OK")

}
