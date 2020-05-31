package StitchingSQLGo

/*
postgres https://www.postgresql.org/docs/current/sql-update.html
*/
type Update struct {
	Table `validate:"required"`
	Set   map[Field]interface{} `validate:"required"`
	Where
	Returning
}

func (u Update) SQL() (string, []interface{}, error) {
	if err := validate.Struct(u); err != nil {
		return "", nil, err
	}

	s := SqlBuilder{}

	// update
	s.WriteString("update")

	// table
	if err := u.Table.Table(&s); err != nil {
		return "", nil, err
	}

	// set
	s.WriteString(" set")

	// field1 = value1 , field2 = value2
	i := 0
	for f, v := range u.Set {
		if err := f.Field(&s, false); err != nil {
			return "", nil, err
		}
		s.WriteString(" =")
		if err := s.push(v); err != nil {
			return "", nil, err
		}

		if i == len(u.Set)-1 {
			break
		}
		s.WriteByte(',')
		i++
	}

	// where field1 = value1 X field2 = value2
	if err := u.Where.where(&s); err != nil {
		return "", nil, err
	}

	return s.String(), s.args, nil
}

func (u Update) Exec() (string, []interface{}, error) {
	return u.SQL()
}

func (u Update) ExecWithReturning() Returning {
	return u.Returning
}
