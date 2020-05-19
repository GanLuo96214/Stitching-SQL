package StitchingSQLGo

type WhereItem interface {
	WhereItem(*sql) error
}

type WhereItems []WhereItem

func (is WhereItems) WhereItem(s *sql) error {
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

func (is WhereItems) WhereItems(s *sql) error {
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
