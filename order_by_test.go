package StitchingSQLGo

import (
	"testing"
)

//go:generate stitching_sql -type=TestOrderByTable -file-name=order_by_test_table_stitching_sql_test.go -is-add-import=false
type TestOrderByTable struct {
	Field1 int
	Field2 int
	Field3 int
}

func TestOrderBy_OrderBy(t *testing.T) {
	s, b := SqlBuilder{}, OrderBy{
		TestOrderByTable_F.Field1: ASC,
		TestOrderByTable_F.Field2: DESC,
		TestOrderByTable_F.Field3: ASC}

	excepts := map[string]interface{}{
		" order by test_order_by_table.field1 asc, test_order_by_table.field2 desc, test_order_by_table.field3 asc": nil,
		" order by test_order_by_table.field2 desc, test_order_by_table.field1 asc, test_order_by_table.field3 asc": nil,
		" order by test_order_by_table.field3 asc, test_order_by_table.field2 desc, test_order_by_table.field1 asc": nil,
		" order by test_order_by_table.field2 desc, test_order_by_table.field3 asc, test_order_by_table.field1 asc": nil,
		" order by test_order_by_table.field3 asc, test_order_by_table.field1 asc, test_order_by_table.field2 desc": nil,
		" order by test_order_by_table.field1 asc, test_order_by_table.field3 asc, test_order_by_table.field2 desc": nil,
	}

	if err := b.orderBy(&s); err != nil {
		t.Fatal(err)
	}

	if _, ok := excepts[s.String()]; ok == false {
		t.Fatalf("now\n%s", s.String())
	}

}
