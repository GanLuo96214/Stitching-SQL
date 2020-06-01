package StitchingSQLGo

import "testing"

//go:generate stitching_sql -type=TestSelect_SQLTable -file-name=sql_select_table_stitching_sql_test.go -is-add-import=false
type TestSelect_SQLTable struct {
	Field1 int
	Field2 int
	Field3 int
}

func TestSelect_SQL(t *testing.T) {

	s := Select{
		Fields: Fields{
			TestSelect_SQLTable_F.Field1,
			TestSelect_SQLTable_F.Field2,
			TestSelect_SQLTable_F.Field3,
		},
		Table: TestSelect_SQLTable{},
		Where: Where{
			Condition{
				L: TestSelect_SQLTable_F.Field1,
				O: EQ,
				R: 1,
			},
			AND,
			WhereItems{
				Condition{
					L: TestSelect_SQLTable_F.Field1,
					O: EQ,
					R: 2,
				},
				AND,
				Condition{
					L: TestSelect_SQLTable_F.Field2,
					O: NEQ,
					R: 3,
				},
				AND,
				Condition{
					L: TestSelect_SQLTable_F.Field2,
					O: GT,
					R: 4,
				},
				AND,
				WhereItems{
					Condition{
						L: TestSelect_SQLTable_F.Field1,
						O: EQ,
						R: 5,
					},
					AND,
					Condition{
						L: TestSelect_SQLTable_F.Field2,
						O: NEQ,
						R: 6,
					},
					AND,
					Condition{
						L: TestSelect_SQLTable_F.Field2,
						O: GT,
						R: 7,
					},
				},
			},
		},
		OrderBy: map[Field]OrderByKind{
			TestSelect_SQLTable_F.Field1: ASC,
			TestSelect_SQLTable_F.Field2: DESC,
			TestSelect_SQLTable_F.Field3: ASC,
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
		"select test_select__sqltable.field1, test_select__sqltable.field2, test_select__sqltable.field3 from test_select__sqltable where test_select__sqltable.field1 = $1 and ( test_select__sqltable.field1 = $2 and test_select__sqltable.field2 != $3 and test_select__sqltable.field2 > $4 and ( test_select__sqltable.field1 = $5 and test_select__sqltable.field2 != $6 and test_select__sqltable.field2 > $7)) order by test_select__sqltable.field1 asc, test_select__sqltable.field2 desc, test_select__sqltable.field3 asc limit 1 offset 1 for update": {1, 2, 3, 4, 5, 6, 7},
		"select test_select__sqltable.field1, test_select__sqltable.field2, test_select__sqltable.field3 from test_select__sqltable where test_select__sqltable.field1 = $1 and ( test_select__sqltable.field1 = $2 and test_select__sqltable.field2 != $3 and test_select__sqltable.field2 > $4 and ( test_select__sqltable.field1 = $5 and test_select__sqltable.field2 != $6 and test_select__sqltable.field2 > $7)) order by test_select__sqltable.field2 desc, test_select__sqltable.field1 asc, test_select__sqltable.field3 asc limit 1 offset 1 for update": {1, 2, 3, 4, 5, 6, 7},
		"select test_select__sqltable.field1, test_select__sqltable.field2, test_select__sqltable.field3 from test_select__sqltable where test_select__sqltable.field1 = $1 and ( test_select__sqltable.field1 = $2 and test_select__sqltable.field2 != $3 and test_select__sqltable.field2 > $4 and ( test_select__sqltable.field1 = $5 and test_select__sqltable.field2 != $6 and test_select__sqltable.field2 > $7)) order by test_select__sqltable.field3 asc, test_select__sqltable.field2 desc, test_select__sqltable.field1 asc limit 1 offset 1 for update": {1, 2, 3, 4, 5, 6, 7},
		"select test_select__sqltable.field1, test_select__sqltable.field2, test_select__sqltable.field3 from test_select__sqltable where test_select__sqltable.field1 = $1 and ( test_select__sqltable.field1 = $2 and test_select__sqltable.field2 != $3 and test_select__sqltable.field2 > $4 and ( test_select__sqltable.field1 = $5 and test_select__sqltable.field2 != $6 and test_select__sqltable.field2 > $7)) order by test_select__sqltable.field2 desc, test_select__sqltable.field3 asc, test_select__sqltable.field1 asc limit 1 offset 1 for update": {1, 2, 3, 4, 5, 6, 7},
		"select test_select__sqltable.field1, test_select__sqltable.field2, test_select__sqltable.field3 from test_select__sqltable where test_select__sqltable.field1 = $1 and ( test_select__sqltable.field1 = $2 and test_select__sqltable.field2 != $3 and test_select__sqltable.field2 > $4 and ( test_select__sqltable.field1 = $5 and test_select__sqltable.field2 != $6 and test_select__sqltable.field2 > $7)) order by test_select__sqltable.field3 asc, test_select__sqltable.field1 asc, test_select__sqltable.field2 desc limit 1 offset 1 for update": {1, 2, 3, 4, 5, 6, 7},
		"select test_select__sqltable.field1, test_select__sqltable.field2, test_select__sqltable.field3 from test_select__sqltable where test_select__sqltable.field1 = $1 and ( test_select__sqltable.field1 = $2 and test_select__sqltable.field2 != $3 and test_select__sqltable.field2 > $4 and ( test_select__sqltable.field1 = $5 and test_select__sqltable.field2 != $6 and test_select__sqltable.field2 > $7)) order by test_select__sqltable.field1 asc, test_select__sqltable.field3 asc, test_select__sqltable.field2 desc limit 1 offset 1 for update": {1, 2, 3, 4, 5, 6, 7},
	}

	exceptArgs, ok := excepts[sql]
	if ok == false {
		t.Fatalf("now\n%s", sql)
	}

	if argsCompare(exceptArgs, args) == false {
		t.Fatalf("except\n%v\nnow\n%v", exceptArgs, args)
	}
}
