package StitchingSQLGo

import "testing"

//go:generate stitching_sql -type=TestWhere_WhereTable -file-name=where_table_stitching_sql_test.go -is-add-import=false
type TestWhere_WhereTable struct {
	Field1 int
	Field2 int
	Field3 int
}

func TestWhere_Where(t *testing.T) {
	s, w := SqlBuilder{}, Where{
		Condition{
			L: TestWhere_WhereTable_F.Field1,
			O: EQ,
			R: 1,
		},
		AND,
		WhereItems{
			Condition{
				L: TestWhere_WhereTable_F.Field1,
				O: EQ,
				R: 2,
			},
			AND,
			Condition{
				L: TestWhere_WhereTable_F.Field2,
				O: NEQ,
				R: 3,
			},
			AND,
			Condition{
				L: TestWhere_WhereTable_F.Field2,
				O: GT,
				R: 4,
			},
			AND,
			WhereItems{
				Condition{
					L: TestWhere_WhereTable_F.Field1,
					O: EQ,
					R: 2,
				},
				AND,
				Condition{
					L: TestWhere_WhereTable_F.Field2,
					O: NEQ,
					R: 3,
				},
				AND,
				Condition{
					L: TestWhere_WhereTable_F.Field2,
					O: GT,
					R: 4,
				},
			},
		},
	}

	if err := w.where(&s); err != nil {
		t.Fatal(err)
	}

	exceptSql := " where test_where__where_table.field1 = $1 and ( test_where__where_table.field1 = $2 and test_where__where_table.field2 != $3 and test_where__where_table.field2 > $4 and ( test_where__where_table.field1 = $5 and test_where__where_table.field2 != $6 and test_where__where_table.field2 > $7))"
	if s.String() != exceptSql {
		t.Fatalf("except\n%s\nnow\n%s", exceptSql, s.String())
	}

	exceptArgs := []interface{}{1, 2, 3, 4, 2, 3, 4}
	if argsCompare(exceptArgs, s.args) == false {
		t.Fatalf("except\n%v\nnow\n%v", exceptArgs, s.args)
	}

}
