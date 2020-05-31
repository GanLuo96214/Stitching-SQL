package StitchingSQLGo

type whereItemLinker string

const (
	AND whereItemLinker = "and" // 1 = 1 and 2 = 2
	OR  whereItemLinker = "or"  // 1 = 1 or  2 = 2
)

func (l whereItemLinker) WhereItem(s *SqlBuilder) error {
	if s == nil {
		return ErrorNilSQL
	}

	s.WriteByte(' ')
	s.WriteString(string(l))

	return nil
}
