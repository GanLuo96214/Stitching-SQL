package StitchingSQLGo

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

func condition_operator(fl validator.FieldLevel) bool {

	conditionOperatorT := reflect.TypeOf((*conditionOperator)(nil)).Elem()

	v := fl.Field()

	if v.Type() != conditionOperatorT {
		return false
	}

	o, ok := v.Interface().(conditionOperator)
	if ok == false {
		return false
	}
	switch o {
	case EQ, NEQ, GT, LT, GTEQ, LTEQ, NGT, NLT, IS, ISNot:
	default:
		return false
	}

	return true
}
