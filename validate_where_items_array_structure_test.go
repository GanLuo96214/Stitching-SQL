package StitchingSQLGo

import (
	"testing"
)

func TestWhereItemArrayStructure(t *testing.T) {
	is := WhereItems{
		Condition{}, AND, Condition{},
	}

	if err := validate.Var(is, "where_items_array_structure"); err != nil {
		t.Fatal(err)
	}

}
