package StitchingSQLGo

func writeSqlTableToStringsBuilder(f SQLTable, s *sql) error {
	if err := validate.Var(f, "required,sql_table_not_return_empty_string"); err != nil {
		return err
	}
	if s == nil {
		return ErrorNilSQL
	}

	s.WriteByte(' ')
	s.WriteString(f.SQLTable())

	return nil
}
