package StitchingSQLGo

import (
	"testing"
)

//go:generate stitching_sql -type=TestCondition_ConditionTable -file-name=condition_table_stitching_sql_test.go -is-add-import=false
type TestCondition_ConditionTable struct {
	Field1 int
}

func TestCondition_Condition(t *testing.T) {
	s, c := SqlBuilder{}, Condition{
		L: TestCondition_ConditionTable_F.Field1,
		O: EQ,
		R: 1,
	}

	if err := c.Condition(&s); err != nil {
		t.Fatal(err)
	}

	sql := " test_condition__condition_table.field1 = $1"
	if s.String() != sql {
		t.Fatalf("except\n%s\nnow\n%s\n", sql, s.String())
	}
	args := []interface{}{1}

	if argsCompare(args, s.args) == false {
		t.Fatalf("except\n%v\nnow\n%v\n", args, s.args)
	}

}
