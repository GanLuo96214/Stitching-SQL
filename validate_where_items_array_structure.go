package StitchingSQLGo

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

func where_items_array_structure(fl validator.FieldLevel) bool {

	whereItemInterfaceT := reflect.TypeOf((*WhereItem)(nil)).Elem()

	v := fl.Field()

	if v.Type().Implements(whereItemInterfaceT) == false {
		return false
	}

	for i := 0; i < v.Len(); i++ {
		eleV := v.Index(i)

		if eleV.IsNil() {
			return false
		}

		_, ok := eleV.Interface().(whereItemLinker)
		if i%2 == 0 && ok == true {
			return false
		} else if i%2 == 1 && ok == false {
			return false
		} else if i == v.Len()-1 && ok == true { // last where_item must not be where_item_linker
			return false
		}
	}

	return true
}
