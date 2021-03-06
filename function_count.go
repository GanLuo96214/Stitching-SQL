package StitchingSQLGo

import "errors"

type Count struct {
	F Field
}

func (e Count) Field(s *SqlBuilder, isRefTable bool) error {
	if s == nil {
		return ErrorNilSQL
	}

	if e.F == nil {
		return errors.New("TODO") // todo
	}

	s.WriteString(" count(")

	if err := e.F.Field(s, isRefTable); err != nil {
		return err
	}

	s.WriteByte(')')
	return nil
}
