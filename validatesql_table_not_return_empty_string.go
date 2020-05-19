package StitchingSQLGo

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

func sql_table_not_return_empty_string(fl validator.FieldLevel) bool {
	sqlTableT := reflect.TypeOf((*SQLTable)(nil)).Elem()

	v := fl.Field()

	if !v.Type().Implements(sqlTableT) {
		return false
	}

	if v := v.MethodByName("SQLTable").Call([]reflect.Value{}); v[0].String() == "" {
		return false
	}

	return true
}
