package StitchingSQLGo

type Where WhereItems // `validate:"required,where_items_array_structure"`

func (w Where) where(s *sql) error {
	if s == nil {
		return ErrorNilSQL
	}
	if w == nil {
		return nil
	}

	s.WriteString(" where")

	return WhereItems(w).WhereItems(s)
}
