package StitchingSQLGo

import "testing"

//go:generate stitching_sql -type=TestDelete_SQLTable -file-name=sql_delete_table_stitching_sql_test.go -is-add-import=false
type TestDelete_SQLTable struct {
	Field1 int
	Field2 int
	Field3 int
}

func TestDelete_SQL(t *testing.T) {
	d := Delete{
		Table: TestDelete_SQLTable{},
		Where: Where{
			Condition{
				L: TestDelete_SQLTable_F.Field1,
				O: EQ,
				R: 1,
			},
			AND,
			WhereItems{
				Condition{
					L: TestDelete_SQLTable_F.Field1,
					O: EQ,
					R: 2,
				},
				AND,
				Condition{
					L: TestDelete_SQLTable_F.Field2,
					O: NEQ,
					R: 3,
				},
				AND,
				Condition{
					L: TestDelete_SQLTable_F.Field2,
					O: GT,
					R: 4,
				},
				AND,
				WhereItems{
					Condition{
						L: TestDelete_SQLTable_F.Field1,
						O: EQ,
						R: 2,
					},
					AND,
					Condition{
						L: TestDelete_SQLTable_F.Field2,
						O: NEQ,
						R: 3,
					},
					AND,
					Condition{
						L: TestDelete_SQLTable_F.Field2,
						O: GT,
						R: 4,
					},
				},
			},
		},
		Returning: Returning{
			TestDelete_SQLTable_F.Field1,
			TestDelete_SQLTable_F.Field2,
			TestDelete_SQLTable_F.Field3,
		},
	}

	sql, args, err := d.SQL()
	if err != nil {
		t.Fatal(err)
	}

	exceptSql := "delete from test_delete__sqltable where test_delete__sqltable.field1 = $1 and ( test_delete__sqltable.field1 = $2 and test_delete__sqltable.field2 != $3 and test_delete__sqltable.field2 > $4 and ( test_delete__sqltable.field1 = $5 and test_delete__sqltable.field2 != $6 and test_delete__sqltable.field2 > $7)) returning test_delete__sqltable.field1, test_delete__sqltable.field2, test_delete__sqltable.field3"
	if sql != exceptSql {
		t.Fatalf("except\n%s\nnow\n%s", exceptSql, sql)
	}

	exceptArgs := []interface{}{1, 2, 3, 4, 2, 3, 4}
	if argsCompare(exceptArgs, args) == false {
		t.Fatalf("except\n%v\nnow\n%v", exceptArgs, args)
	}

}
