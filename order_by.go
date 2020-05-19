package StitchingSQLGo

type OrderByKind string

const (
	ASC  OrderByKind = "asc"
	DESC OrderByKind = "desc"
)

type OrderBy map[SQLField]OrderByKind

func (b OrderBy) orderBy(s *sql) error {
	if s == nil {
		return ErrorNilSQL
	}
	if b == nil {
		return nil
	}

	if err := validate.Var(b, "required,dive,keys,required,endkeys,eq=asc|eq=desc"); err != nil {
		return err
	}

	s.WriteString(" order by")

	i := 0
	for f, k := range b {
		if err := WriteSqlFieldToStringsBuilder(f, s, true); err != nil {
			return err
		}
		s.WriteByte(' ')
		s.WriteString(string(k))

		if i == len(b)-1 {
			break
		}
		s.WriteByte(',')
		i++
	}

	return nil
}
