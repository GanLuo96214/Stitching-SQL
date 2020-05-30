package StitchingSQLGo

import (
	"errors"
)

type conditionOperator string

const (
	EQ    conditionOperator = "="      // Equal (=) Operator
	NEQ   conditionOperator = "!="     // Not Equal (!= or <>) Operator
	GT    conditionOperator = ">"      // Greater Than (>) Operator
	LT    conditionOperator = "<"      // Less Than (<) Operator
	GTEQ  conditionOperator = ">="     // Greater Than or Equal To (>=) Operator
	LTEQ  conditionOperator = "<="     // Less Than or Equal To (<=) Operator
	NGT   conditionOperator = "!>"     // Not Greater Than (!>) Operator
	NLT   conditionOperator = "!<"     // Not Less Than (!<) Operator
	IS    conditionOperator = "is"     // IS Operator
	ISNot conditionOperator = "is not" // IS NOT Operator
	//In     conditionOperator = "in"      // In Operator
	//NotInt conditionOperator = "not int" // NOT IN Operator
)

var ErrorNotConditionOperator = errors.New("not a condition operator")

func (o conditionOperator) conditionOperator(s *sql) error {
	if s == nil {
		return ErrorNilSQL
	}
	if err := validate.Var(o, "condition_operator"); err != nil {
		return err
	}

	s.WriteByte(' ')
	s.WriteString(string(o))

	return nil
}
