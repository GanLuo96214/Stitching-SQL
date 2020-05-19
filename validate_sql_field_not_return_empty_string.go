package StitchingSQLGo

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func sql_field_not_return_empty_string(fl validator.FieldLevel) bool {

	sqlTableT := reflect.TypeOf((*SQLTable)(nil)).Elem()
	sqlFieldT := reflect.TypeOf((*SQLField)(nil)).Elem()

	v := fl.Field()

	if v.Type().Implements(sqlTableT) == false || v.Type().Implements(sqlFieldT) == false {
		return false
	}

	if v := v.MethodByName("SQLTable").Call([]reflect.Value{}); strings.TrimSpace(v[0].String()) == "" {
		return false
	}

	if v := v.MethodByName("SQLField").Call([]reflect.Value{}); strings.TrimSpace(v[0].String()) == "" {
		return false
	}

	return true
}
