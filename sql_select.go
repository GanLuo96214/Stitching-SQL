package StitchingSQLGo

/*
postgres https://www.postgresql.org/docs/current/sql-select.html
*/
type Select struct {
	SQLFields `validate:"required"`
	SQLTable  `validate:"required,sql_table_not_return_empty_string"`
	Where
	OrderBy
	Limit
	Offset
	For
}

func (slt Select) SQL() (string, []interface{}, error) {
	if err := validate.Struct(slt); err != nil {
		return "", nil, err
	}

	s := sql{}

	// select
	s.WriteString("select")

	// field1, field2
	if err := slt.SQLFields.sqlFields(&s); err != nil {
		return "", nil, err
	}

	// from
	s.WriteString(" from")

	// table
	if err := writeSqlTableToStringsBuilder(slt.SQLTable, &s); err != nil {
		return "", nil, err
	}

	// where field1 = value1 and filed2 = value2
	if err := slt.Where.where(&s); err != nil {
		return "", nil, err
	}

	// order by field1 asc, field2 desc
	if err := slt.OrderBy.orderBy(&s); err != nil {
		return "", nil, err
	}

	// limit X
	if err := slt.Limit.limit(&s); err != nil {
		return "", nil, err
	}

	// offset X
	if err := slt.Offset.offset(&s); err != nil {
		return "", nil, err
	}

	// for update
	if err := slt.For.For(&s); err != nil {
		return "", nil, err
	}

	return s.String(), s.args, nil
}
