package StitchingSQLGo

/*
postgres https://www.postgresql.org/docs/current/sql-select.html
*/
type SelectExists struct {
	Fields `validate:"required"`
	Table  `validate:"required"`
	Where
	OrderBy
	Limit
	Offset
	For
}

func (slt SelectExists) SQL() (string, []interface{}, error) {
	if err := validate.Struct(slt); err != nil {
		return "", nil, err
	}

	s := SqlBuilder{}

	// select exists(select
	s.WriteString("select exists(select")

	// field1, field2
	if err := slt.Fields.Fields(&s); err != nil {
		return "", nil, err
	}

	// from
	s.WriteString(" from")

	// table
	if err := slt.Table.Table(&s); err != nil {
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

	// )
	s.WriteByte(')')

	// for update
	if err := slt.For.For(&s); err != nil {
		return "", nil, err
	}

	return s.String(), s.args, nil
}
