package StitchingSQLGo

type WhereItem interface {
	WhereItem(*SqlBuilder) error
}

type WhereItems []WhereItem

func (is WhereItems) WhereItem(s *SqlBuilder) error {
	if s == nil {
		return ErrorNilSQL
	}

	if err := validate.Var(is, "where_items_array_structure"); err != nil {
		return err
	}

	s.WriteString(" (")
	for _, i := range is {
		if err := i.WhereItem(s); err != nil {
			return err
		}
	}
	s.WriteByte(')')

	return nil
}

func (is WhereItems) WhereItems(s *SqlBuilder) error {
	if s == nil {
		return ErrorNilSQL
	}

	if err := validate.Var(is, "where_items_array_structure"); err != nil {
		return err
	}

	for _, i := range is {
		if err := i.WhereItem(s); err != nil {
			return err
		}
	}
	return nil
}
