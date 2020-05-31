package StitchingSQLGo

import "testing"

//go:generate stitching_sql -type=TestWhereItems_WhereItemTable -file-name=where_items_table_stitching_sql_test.go -is-add-import=false
type TestWhereItems_WhereItemTable struct {
	Field1 int
	Field2 int
	Field3 int
}

func TestWhereItems_WhereItem(t *testing.T) {
	s, is := SqlBuilder{}, WhereItems{
		Condition{
			L: TestWhereItems_WhereItemTable_F.Field1,
			O: EQ,
			R: 1,
		},
		AND,
		WhereItems{
			Condition{
				L: TestWhereItems_WhereItemTable_F.Field1,
				O: EQ,
				R: 2,
			},
			AND,
			Condition{
				L: TestWhereItems_WhereItemTable_F.Field2,
				O: NEQ,
				R: 3,
			},
			AND,
			Condition{
				L: TestWhereItems_WhereItemTable_F.Field2,
				O: GT,
				R: 4,
			},
			AND,
			WhereItems{
				Condition{
					L: TestWhereItems_WhereItemTable_F.Field1,
					O: EQ,
					R: 2,
				},
				AND,
				Condition{
					L: TestWhereItems_WhereItemTable_F.Field2,
					O: NEQ,
					R: 3,
				},
				AND,
				Condition{
					L: TestWhereItems_WhereItemTable_F.Field2,
					O: GT,
					R: 4,
				},
			},
		},
	}

	if err := is.WhereItems(&s); err != nil {
		t.Fatal(err)
	}

	exceptSql := " test_where_items__where_item_table.field1 = $1 and ( test_where_items__where_item_table.field1 = $2 and test_where_items__where_item_table.field2 != $3 and test_where_items__where_item_table.field2 > $4 and ( test_where_items__where_item_table.field1 = $5 and test_where_items__where_item_table.field2 != $6 and test_where_items__where_item_table.field2 > $7))"
	if s.String() != exceptSql {
		t.Fatalf("except\n%s\nnow\n%s", exceptSql, s.String())
	}

	exceptArgs := []interface{}{
		1, 2, 3, 4, 2, 3, 4,
	}
	if argsCompare(exceptArgs, s.args) == false {
		t.Fatalf("except\n%v\nnow\n%v", exceptArgs, s.args)
	}

}
