package StitchingSQLGo

import "testing"

type TestTestValidateSqlTableNotEmptyStringTable byte

func (t TestTestValidateSqlTableNotEmptyStringTable) SQLTable() string {
	return "1"
}

func TestValidateSqlTableNotEmptyString(t *testing.T) {
	table := TestTestValidateSqlTableNotEmptyStringTable(0)
	if err := validate.Var(table, "sql_table_not_return_empty_string"); err != nil {
		t.Fatal(err)
	}
}
