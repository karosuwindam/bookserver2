package sqlite

import "testing"

func TestOpen(t *testing.T) {
	sql := Setup("test.db")
	err := sql.Open()
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	err = sql.Close()
	if err != nil {
		t.Errorf("%v", err.Error())

	}
}
