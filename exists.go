package StitchingSQLGo

import "errors"

type Exists struct {
	F Field
}

func (e Exists) Field(s *SqlBuilder, isRefTable bool) error {
	if s == nil {
		return ErrorNilSQL
	}

	if e.F == nil {
		return errors.New("TODO") // todo
	}

	s.WriteString(" exists(")

	if err := e.F.Field(s, isRefTable); err != nil {
		return err
	}

	s.WriteByte(')')
	return nil
}
