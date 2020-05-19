package StitchingSQLGo

import (
	"github.com/go-playground/validator/v10"
	"log"
)

var (
	validate = validator.New()
)

func init() {

	if err := validate.RegisterValidation("condition_operator", condition_operator); err != nil {
		log.Fatal(err)
	}

	if err := validate.RegisterValidation("sql_table_not_return_empty_string", sql_table_not_return_empty_string); err != nil {
		log.Fatal(err)
	}

	if err := validate.RegisterValidation("sql_field_not_return_empty_string", sql_field_not_return_empty_string); err != nil {
		log.Fatal(err)
	}

	if err := validate.RegisterValidation("where_items_array_structure", where_items_array_structure); err != nil {
		log.Fatal(err)
	}

}
