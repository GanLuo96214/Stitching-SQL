package StitchingSQLGo

import "testing"

//go:generate stitching_sql -type=TestSelectExists_SQLTable -file-name=sql_select_exists_table_stitching_sql_test.go -is-add-import=false
type TestSelectExists_SQLTable struct {
	Field1 int
	Field2 int
	Field3 int
}

func TestSelectExists_SQL(t *testing.T) {

	s := SelectExists{
		FS: Fields{
			Field1{},
			TestSelectExists_SQLTable_F.Field1,
			TestSelectExists_SQLTable_F.Field2,
			TestSelectExists_SQLTable_F.Field3,
			Field1{},
		},
		Table: TestSelectExists_SQLTable{},
		Where: Where{
			Condition{
				L: TestSelectExists_SQLTable_F.Field1,
				O: EQ,
				R: 1,
			},
			AND,
			WhereItems{
				Condition{
					L: TestSelectExists_SQLTable_F.Field1,
					O: EQ,
					R: 2,
				},
				AND,
				Condition{
					L: TestSelectExists_SQLTable_F.Field2,
					O: NEQ,
					R: 3,
				},
				AND,
				Condition{
					L: TestSelectExists_SQLTable_F.Field2,
					O: GT,
					R: 4,
				},
				AND,
				WhereItems{
					Condition{
						L: TestSelectExists_SQLTable_F.Field1,
						O: EQ,
						R: 5,
					},
					AND,
					Condition{
						L: TestSelectExists_SQLTable_F.Field2,
						O: NEQ,
						R: 6,
					},
					AND,
					Condition{
						L: TestSelectExists_SQLTable_F.Field2,
						O: GT,
						R: 7,
					},
				},
			},
		},
		OrderBy: map[Field]OrderByKind{
			TestSelectExists_SQLTable_F.Field1: ASC,
			TestSelectExists_SQLTable_F.Field2: DESC,
			TestSelectExists_SQLTable_F.Field3: ASC,
		},
		Limit:  1,
		Offset: 1,
		For:    ForUpdate,
	}

	sql, args, err := s.SQL()
	if err != nil {
		t.Fatal(err)
	}

	excepts := map[string][]interface{}{
		"select exists(select 1, test_select_exists__sqltable.field1, test_select_exists__sqltable.field2, test_select_exists__sqltable.field3, 1 from test_select_exists__sqltable where test_select_exists__sqltable.field1 = $1 and ( test_select_exists__sqltable.field1 = $2 and test_select_exists__sqltable.field2 != $3 and test_select_exists__sqltable.field2 > $4 and ( test_select_exists__sqltable.field1 = $5 and test_select_exists__sqltable.field2 != $6 and test_select_exists__sqltable.field2 > $7)) order by test_select_exists__sqltable.field1 asc, test_select_exists__sqltable.field2 desc, test_select_exists__sqltable.field3 asc limit 1 offset 1) for update": {1, 2, 3, 4, 5, 6, 7},
		"select exists(select 1, test_select_exists__sqltable.field1, test_select_exists__sqltable.field2, test_select_exists__sqltable.field3, 1 from test_select_exists__sqltable where test_select_exists__sqltable.field1 = $1 and ( test_select_exists__sqltable.field1 = $2 and test_select_exists__sqltable.field2 != $3 and test_select_exists__sqltable.field2 > $4 and ( test_select_exists__sqltable.field1 = $5 and test_select_exists__sqltable.field2 != $6 and test_select_exists__sqltable.field2 > $7)) order by test_select_exists__sqltable.field2 desc, test_select_exists__sqltable.field1 asc, test_select_exists__sqltable.field3 asc limit 1 offset 1) for update": {1, 2, 3, 4, 5, 6, 7},
		"select exists(select 1, test_select_exists__sqltable.field1, test_select_exists__sqltable.field2, test_select_exists__sqltable.field3, 1 from test_select_exists__sqltable where test_select_exists__sqltable.field1 = $1 and ( test_select_exists__sqltable.field1 = $2 and test_select_exists__sqltable.field2 != $3 and test_select_exists__sqltable.field2 > $4 and ( test_select_exists__sqltable.field1 = $5 and test_select_exists__sqltable.field2 != $6 and test_select_exists__sqltable.field2 > $7)) order by test_select_exists__sqltable.field3 asc, test_select_exists__sqltable.field2 desc, test_select_exists__sqltable.field1 asc limit 1 offset 1) for update": {1, 2, 3, 4, 5, 6, 7},
		"select exists(select 1, test_select_exists__sqltable.field1, test_select_exists__sqltable.field2, test_select_exists__sqltable.field3, 1 from test_select_exists__sqltable where test_select_exists__sqltable.field1 = $1 and ( test_select_exists__sqltable.field1 = $2 and test_select_exists__sqltable.field2 != $3 and test_select_exists__sqltable.field2 > $4 and ( test_select_exists__sqltable.field1 = $5 and test_select_exists__sqltable.field2 != $6 and test_select_exists__sqltable.field2 > $7)) order by test_select_exists__sqltable.field2 desc, test_select_exists__sqltable.field3 asc, test_select_exists__sqltable.field1 asc limit 1 offset 1) for update": {1, 2, 3, 4, 5, 6, 7},
		"select exists(select 1, test_select_exists__sqltable.field1, test_select_exists__sqltable.field2, test_select_exists__sqltable.field3, 1 from test_select_exists__sqltable where test_select_exists__sqltable.field1 = $1 and ( test_select_exists__sqltable.field1 = $2 and test_select_exists__sqltable.field2 != $3 and test_select_exists__sqltable.field2 > $4 and ( test_select_exists__sqltable.field1 = $5 and test_select_exists__sqltable.field2 != $6 and test_select_exists__sqltable.field2 > $7)) order by test_select_exists__sqltable.field3 asc, test_select_exists__sqltable.field1 asc, test_select_exists__sqltable.field2 desc limit 1 offset 1) for update": {1, 2, 3, 4, 5, 6, 7},
		"select exists(select 1, test_select_exists__sqltable.field1, test_select_exists__sqltable.field2, test_select_exists__sqltable.field3, 1 from test_select_exists__sqltable where test_select_exists__sqltable.field1 = $1 and ( test_select_exists__sqltable.field1 = $2 and test_select_exists__sqltable.field2 != $3 and test_select_exists__sqltable.field2 > $4 and ( test_select_exists__sqltable.field1 = $5 and test_select_exists__sqltable.field2 != $6 and test_select_exists__sqltable.field2 > $7)) order by test_select_exists__sqltable.field1 asc, test_select_exists__sqltable.field3 asc, test_select_exists__sqltable.field2 desc limit 1 offset 1) for update": {1, 2, 3, 4, 5, 6, 7},
	}

	exceptArgs, ok := excepts[sql]
	if ok == false {
		t.Fatalf("now\n%s", sql)
	}

	if argsCompare(exceptArgs, args) == false {
		t.Fatalf("except\n%v\nnow\n%v", exceptArgs, args)
	}
}
