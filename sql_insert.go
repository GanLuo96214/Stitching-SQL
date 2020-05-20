package StitchingSQLGo

/*
postgres https://www.postgresql.org/docs/current/sql-insert.html
*/
type Insert struct {
	SQLTable `validate:"required,sql_table_not_return_empty_string"`
	Values   map[SQLField]interface{} `validate:"required,dive,keys,required,endkeys,required"`
	Returning
}

func (i Insert) SQL() (string, []interface{}, error) {
	if err := validate.Struct(i); err != nil {
		return "", nil, err
	}

	s := sql{}

	// insert into
	s.WriteString("insert into")

	// table
	if err := writeSqlTableToStringsBuilder(i.SQLTable, &s); err != nil {
		return "", nil, err
	}

	// (field1,fiedl2)
	keys := make([]SQLField, 0, len(i.Values))
	s.WriteString(" (")
	for f := range i.Values {
		if err := WriteSqlFieldToStringsBuilder(f, &s, false); err != nil {
			return "", nil, err
		}
		keys = append(keys, f)

		if len(keys) == len(i.Values) {
			break
		}
		s.WriteByte(',')
	}
	s.WriteByte(')')

	// values
	s.WriteString(" values")

	// (value1,value2)
	s.WriteString(" (")
	for idx, f := range keys {
		if err := s.push(i.Values[f]); err != nil {
			return "", nil, err
		}

		if idx == len(keys)-1 {
			break
		}
		s.WriteByte(',')
	}
	s.WriteByte(')')

	// returning
	if err := i.Returning.Returning(&s); err != nil {
		return "", nil, err
	}

	return s.String(), s.args, nil
}

func (i Insert) Exec() (string, []interface{}, error) {
	return i.SQL()
}
