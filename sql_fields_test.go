package StitchingSQLGo

import "testing"

//go:generate stitching_sql -type=TestSQLFields_SQLFieldsTable -file-name=sql_fields_table_stitching_sql_test.go -is-add-import=false
type TestSQLFields_SQLFieldsTable struct {
	Field1 int
	Field2 int
	Field3 int
}

func TestSQLFields_SQLFields(t *testing.T) {
	s, fs := sql{}, SQLFields{
		TestSQLFields_SQLFieldsTable_F.Field1,
		TestSQLFields_SQLFieldsTable_F.Field2,
		TestSQLFields_SQLFieldsTable_F.Field3,
	}

	if err := fs.sqlFields(&s); err != nil {
		t.Fatal(err)
	}

	except := " test_sqlfields__sqlfields_table.field1, test_sqlfields__sqlfields_table.field2, test_sqlfields__sqlfields_table.field3"
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}

}
