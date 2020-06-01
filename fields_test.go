package StitchingSQLGo

import "testing"

//go:generate stitching_sql -type=TestFields_FieldsTable -file-name=fields_table_stitching_sql_test.go -is-add-import=false
type TestFields_FieldsTable struct {
	Field1 int
	Field2 int
	Field3 int
}

func TestFields_Fields(t *testing.T) {
	s, fs := SqlBuilder{}, Fields{
		TestFields_FieldsTable_F.Field1,
		TestFields_FieldsTable_F.Field2,
		TestFields_FieldsTable_F.Field3,
	}

	if err := fs.Fields(&s); err != nil {
		t.Fatal(err)
	}

	except := " test_fields__fields_table.field1, test_fields__fields_table.field2, test_fields__fields_table.field3"
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}

}
