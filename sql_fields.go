package StitchingSQLGo

type SQLFields []SQLField

func (fs SQLFields) sqlFields(s *sql) error {
	if err := validate.Var(fs, "required,dive,required"); err != nil {
		return err
	}

	for i, f := range fs {
		if err := WriteSqlFieldToStringsBuilder(f, s, true); err != nil {
			return err
		}
		if i == len(fs)-1 {
			break
		}
		s.WriteByte(',')
	}

	return nil
}
