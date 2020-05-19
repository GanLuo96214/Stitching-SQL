package StitchingSQLGo

/*
postgres https://www.postgresql.org/docs/current/sql-delete.html
*/
type Delete struct {
	SQLTable `validate:"required,sql_table_not_return_empty_string"`
	Where    `validate:"required"`
	Returning
}

func (d Delete) SQL() (string, []interface{}, error) {
	if err := validate.Struct(d); err != nil {
		return "", nil, err
	}

	s := sql{}

	s.WriteString("delete from")
	if err := writeSqlTableToStringsBuilder(d.SQLTable, &s); err != nil {
		return "", nil, err
	}

	if err := d.Where.where(&s); err != nil {
		return "", nil, err
	}

	if err := d.Returning.Returning(&s); err != nil {
		return "", nil, err
	}

	return s.String(), s.args, nil
}
