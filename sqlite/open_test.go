package sqlite

import (
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	testdbname := "test.db"

	sql := Setup(testdbname)
	err := sql.Open()
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	err = sql.Close()
	if err != nil {
		t.Errorf("%v", err.Error())

	}
	os.Remove(testdbname)
}
