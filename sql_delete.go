package StitchingSQLGo

/*
postgres https://www.postgresql.org/docs/current/sql-delete.html
*/
type Delete struct {
	Table `validate:"required"`
	Where `validate:"required"`
	Returning
}

func (d Delete) SQL() (string, []interface{}, error) {
	if err := validate.Struct(d); err != nil {
		return "", nil, err
	}

	s := SqlBuilder{}

	s.WriteString("delete from")
	if err := d.Table.Table(&s); err != nil {
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

func (d Delete) Exec() (string, []interface{}, error) {
	return d.SQL()
}

func (d Delete) ExecWithReturning() Returning {
	return d.Returning
}
