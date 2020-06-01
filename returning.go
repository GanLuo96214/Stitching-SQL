package StitchingSQLGo

type Returning Fields

func (r Returning) Returning(s *SqlBuilder) error {
	if s == nil {
		return ErrorNilSQL
	}
	if r == nil {
		return nil
	}

	s.WriteString(" returning")

	return Fields(r).Fields(s)
}
