package StitchingSQLGo

import (
	"testing"
)

//go:generate stitching_sql -type=TestWriteSqlTableToStringsBuilderTable -file-name=write_sql_table_to_strings_builder_table_stitching_sql_test.go -is-add-import=false
type TestWriteSqlTableToStringsBuilderTable struct{}

func TestWriteSqlTableToStringsBuilder(t *testing.T) {
	s := sql{}

	except := " test_write_sql_table_to_strings_builder_table"
	if err := writeSqlTableToStringsBuilder(TestWriteSqlTableToStringsBuilderTable{}, &s); err != nil {
		t.Fatal(err)
	}

	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}
}
