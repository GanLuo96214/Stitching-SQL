package StitchingSQLGo

type Field1 struct{}

func (f Field1) Field(s *SqlBuilder, isRefTable bool) error {
	if s == nil {
		return ErrorNilSQL
	}

	s.WriteByte('1')

	return nil
}
