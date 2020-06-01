package StitchingSQLGo

import (
	"testing"
)

//go:generate stitching_sql -type=TestExists_FieldTable -file-name=function_exists_table_stitching_sql_test.go -is-add-import=false
type TestExists_FieldTable struct {
	Field1 int
}

func TestExists_Field_RefTableTrue(t *testing.T) {
	s, e := SqlBuilder{}, Exists{
		F: TestExists_FieldTable_F.Field1,
	}

	if err := e.Field(&s, true); err != nil {
		t.Fatal(err)
	}

	except := " exists( test_exists__field_table.field1)"
	if except != s.String() {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}

}

func TestExists_Field_RefTableFalse(t *testing.T) {
	s, e := SqlBuilder{}, Exists{
		F: TestExists_FieldTable_F.Field1,
	}

	if err := e.Field(&s, false); err != nil {
		t.Fatal(err)
	}

	except := " exists( field1)"
	if except != s.String() {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}

}
