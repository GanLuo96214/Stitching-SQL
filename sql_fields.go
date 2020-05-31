package StitchingSQLGo

type SQLFields []Field

func (fs SQLFields) sqlFields(s *SqlBuilder) error {
	if s == nil {
		return ErrorNilSQL
	}
	if err := validate.Var(fs, "required,dive,required"); err != nil {
		return err
	}

	for i, f := range fs {
		if err := f.Field(s, true); err != nil {
			return err
		}
		if i == len(fs)-1 {
			break
		}
		s.WriteByte(',')
	}

	return nil
}
