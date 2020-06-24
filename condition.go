package StitchingSQLGo

/*
condition

L field
O condition operator
R value
*/
type Condition struct {
	L Field             `validate:"required"`
	O conditionOperator `validate:"required"`
	R interface{}       `validate:"required"`
}

func (c Condition) Condition(s *SqlBuilder) error {
	if s == nil {
		return ErrorNilSQL
	}

	if err := validate.Struct(c); err != nil {
		return err
	}

	// c.L
	if err := c.L.Field(s, true); err != nil {
		return err
	}

	// c.O
	if err := c.O.conditionOperator(s); err != nil {
		return err
	}

	// c.R
	if err := s.push(c.R); err != nil {
		return err
	}

	return nil
}

func (c Condition) WhereItem(s *SqlBuilder) error {
	if s == nil {
		return ErrorNilSQL
	}
	return c.Condition(s)
}
