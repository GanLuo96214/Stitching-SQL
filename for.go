package StitchingSQLGo

type For string

const (
	ForUpdate For = "update"
)

func (f For) For(s *SqlBuilder) error {
	if s == nil {
		return ErrorNilSQL
	}
	if f == "" {
		return nil
	}

	s.WriteString(" for ")
	s.WriteString(string(f))

	return nil
}
