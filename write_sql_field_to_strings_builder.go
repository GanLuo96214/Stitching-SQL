package StitchingSQLGo

func WriteSqlFieldToStringsBuilder(f SQLField, s *sql, isRefTable bool) error {
	if s == nil {
		return ErrorNilSQL
	}
	if err := validate.Var(f, "required,sql_field_not_return_empty_string"); err != nil {
		return err
	}

	s.WriteByte(' ')

	if isRefTable {
		s.WriteString(f.SQLTable())
		s.WriteByte('.')
	}

	s.WriteString(f.SQLField())

	return nil
}
