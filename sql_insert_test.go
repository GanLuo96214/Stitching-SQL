package StitchingSQLGo

import (
	"testing"
)

//go:generate stitching_sql -type=TestInsert_SQLTable -file-name=sql_insert_table_stitching_sql_test.go -is-add-import=false
type TestInsert_SQLTable struct {
	Field1 int
	Field2 int
	Field3 int
}

func TestInsert_SQL(t *testing.T) {
	i := Insert{
		Table: TestInsert_SQLTable{},
		Values: map[Field]interface{}{
			TestInsert_SQLTable_F.Field1: 1,
			TestInsert_SQLTable_F.Field2: 2,
			TestInsert_SQLTable_F.Field3: 3,
		},
		Returning: Returning{
			TestInsert_SQLTable_F.Field1,
			TestInsert_SQLTable_F.Field2,
			TestInsert_SQLTable_F.Field3,
		},
	}

	sql, args, err := i.SQL()
	if err != nil {
		t.Fatal(err)
	}

	a := map[string][]interface{}{
		"insert into test_insert__sqltable ( field1, field2, field3) values ( $1, $2, $3) returning test_insert__sqltable.field1, test_insert__sqltable.field2, test_insert__sqltable.field3": {1, 2, 3},
		"insert into test_insert__sqltable ( field2, field1, field3) values ( $1, $2, $3) returning test_insert__sqltable.field1, test_insert__sqltable.field2, test_insert__sqltable.field3": {2, 1, 3},
		"insert into test_insert__sqltable ( field3, field2, field1) values ( $1, $2, $3) returning test_insert__sqltable.field1, test_insert__sqltable.field2, test_insert__sqltable.field3": {3, 2, 1},
		"insert into test_insert__sqltable ( field2, field3, field1) values ( $1, $2, $3) returning test_insert__sqltable.field1, test_insert__sqltable.field2, test_insert__sqltable.field3": {2, 3, 1},
		"insert into test_insert__sqltable ( field3, field1, field2) values ( $1, $2, $3) returning test_insert__sqltable.field1, test_insert__sqltable.field2, test_insert__sqltable.field3": {3, 1, 2},
		"insert into test_insert__sqltable ( field1, field3, field2) values ( $1, $2, $3) returning test_insert__sqltable.field1, test_insert__sqltable.field2, test_insert__sqltable.field3": {1, 3, 2},
	}

	exceptArgs, ok := a[sql]
	if ok == false {
		t.Fatalf("now\n%s", sql)
	}

	if argsCompare(exceptArgs, args) == false {
		t.Fatalf("except\n%v\nnow\n%v", exceptArgs, args)
	}

}
