package StitchingSQLGo

import "testing"

type TestValidateSqlFieldNotEmptyStringTable struct{}

func (t TestValidateSqlFieldNotEmptyStringTable) SQLTable() string {
	return "1"
}
func (t TestValidateSqlFieldNotEmptyStringTable) SQLField() string {
	return "2"
}

func TestValidateSqlFieldNotReturnEmptyString(t *testing.T) {
	table := TestValidateSqlFieldNotEmptyStringTable{}
	if err := validate.Var(table, "sql_field_not_return_empty_string"); err != nil {
		t.Fatal(err)
	}
}
