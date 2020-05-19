package StitchingSQLGo

import "testing"

//go:generate stitching_sql -type=TestReturning_ReturningTable -file-name=returning_table_stitching_sql_test.go -is-add-import=false
type TestReturning_ReturningTable struct {
	Field1 int
	Field2 int
	Field3 int
}

func TestReturning_Returning(t *testing.T) {

	s, r := sql{}, Returning{
		TestReturning_ReturningTable_F.Field1,
		TestReturning_ReturningTable_F.Field2,
		TestReturning_ReturningTable_F.Field3,
	}

	if err := r.Returning(&s); err != nil {
		t.Fatal(err)
	}

	except := " returning test_returning__returning_table.field1, test_returning__returning_table.field2, test_returning__returning_table.field3"
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}

}
