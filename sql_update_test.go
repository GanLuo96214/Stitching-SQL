package StitchingSQLGo

import "testing"

//go:generate stitching_sql -type=TestUpdate_SQLTable -file-name=sql_update_table_stitching_sql_test.go -is-add-import=false
type TestUpdate_SQLTable struct {
	Field1 int
	Field2 int
	Field3 int
}

func TestUpdate_SQL(t *testing.T) {
	u := Update{
		SQLTable: TestUpdate_SQLTable{},
		Set: map[SQLField]interface{}{
			TestUpdate_SQLTable_F.Field1: 1,
			TestUpdate_SQLTable_F.Field2: 2,
			TestUpdate_SQLTable_F.Field3: 3,
		},
		Where: Where{
			Condition{
				L: TestUpdate_SQLTable_F.Field1,
				O: EQ,
				R: 4,
			},
			AND,
			WhereItems{
				Condition{
					L: TestUpdate_SQLTable_F.Field1,
					O: EQ,
					R: 5,
				},
				AND,
				Condition{
					L: TestUpdate_SQLTable_F.Field2,
					O: EQ,
					R: 6,
				},
				AND,
				Condition{
					L: TestUpdate_SQLTable_F.Field3,
					O: EQ,
					R: 7,
				},
			},
		},
		Returning: Returning{
			TestUpdate_SQLTable_F.Field1,
			TestUpdate_SQLTable_F.Field2,
			TestUpdate_SQLTable_F.Field3,
		},
	}

	sql, args, err := u.SQL()
	if err != nil {
		t.Fatal(err)
	}

	a := map[string][]interface{}{
		"update test_update__sqltable set field1 = $1, field2 = $2, field3 = $3 where test_update__sqltable.field1 = $4 and ( test_update__sqltable.field1 = $5 and test_update__sqltable.field2 = $6 and test_update__sqltable.field3 = $7)": {1, 2, 3, 4, 5, 6, 7},
		"update test_update__sqltable set field2 = $1, field1 = $2, field3 = $3 where test_update__sqltable.field1 = $4 and ( test_update__sqltable.field1 = $5 and test_update__sqltable.field2 = $6 and test_update__sqltable.field3 = $7)": {2, 1, 3, 4, 5, 6, 7},
		"update test_update__sqltable set field3 = $1, field2 = $2, field1 = $3 where test_update__sqltable.field1 = $4 and ( test_update__sqltable.field1 = $5 and test_update__sqltable.field2 = $6 and test_update__sqltable.field3 = $7)": {3, 2, 1, 4, 5, 6, 7},
		"update test_update__sqltable set field2 = $1, field3 = $2, field1 = $3 where test_update__sqltable.field1 = $4 and ( test_update__sqltable.field1 = $5 and test_update__sqltable.field2 = $6 and test_update__sqltable.field3 = $7)": {2, 3, 1, 4, 5, 6, 7},
		"update test_update__sqltable set field3 = $1, field1 = $2, field2 = $3 where test_update__sqltable.field1 = $4 and ( test_update__sqltable.field1 = $5 and test_update__sqltable.field2 = $6 and test_update__sqltable.field3 = $7)": {3, 1, 2, 4, 5, 6, 7},
		"update test_update__sqltable set field1 = $1, field3 = $2, field2 = $3 where test_update__sqltable.field1 = $4 and ( test_update__sqltable.field1 = $5 and test_update__sqltable.field2 = $6 and test_update__sqltable.field3 = $7)": {1, 3, 2, 4, 5, 6, 7},
	}

	exceptArgs, ok := a[sql]
	if ok == false {
		t.Fatalf("now\n%s", sql)
	}

	if argsCompare(exceptArgs, args) == false {
		t.Fatalf("except\n%v\nnow\n%v", exceptArgs, args)
	}

}
