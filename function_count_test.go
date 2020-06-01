package StitchingSQLGo

import (
	"testing"
)

//go:generate stitching_sql -type=TestCount_FieldTable -file-name=function_Count_table_stitching_sql_test.go -is-add-import=false
type TestCount_FieldTable struct {
	Field1 int
}

func TestCount_Field_RefTableTrue(t *testing.T) {
	s, e := SqlBuilder{}, Count{
		F: TestCount_FieldTable_F.Field1,
	}

	if err := e.Field(&s, true); err != nil {
		t.Fatal(err)
	}

	except := " count( test_count__field_table.field1)"
	if except != s.String() {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}

}

func TestCount_Field_RefTableFalse(t *testing.T) {
	s, e := SqlBuilder{}, Count{
		F: TestCount_FieldTable_F.Field1,
	}

	if err := e.Field(&s, false); err != nil {
		t.Fatal(err)
	}

	except := " count( field1)"
	if except != s.String() {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}

}

func TestCount1(t *testing.T) {
	s, e := SqlBuilder{}, Count{
		F: Count1{},
	}

	if err := e.Field(&s, false); err != nil {
		t.Fatal(err)
	}

	except := " count(1)"
	if except != s.String() {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}

}
