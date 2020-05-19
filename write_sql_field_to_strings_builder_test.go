package StitchingSQLGo

import (
	"testing"
)

//go:generate stitching_sql -type=TestWriteSqlFieldToStringsBuilderTable -file-name=write_sql_field_to_strings_builder_table_stitching_sql_test.go -is-add-import=false
type TestWriteSqlFieldToStringsBuilderTable struct {
	Field1 int
}

func TestWriteSqlFieldToStringsBuilder(t *testing.T) {
	s := sql{}
	except := " test_write_sql_field_to_strings_builder_table.field1"
	if err := WriteSqlFieldToStringsBuilder(TestWriteSqlFieldToStringsBuilderTable_F.Field1, &s, true); err != nil {
		t.Fatal(err)
	} else if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}
}
