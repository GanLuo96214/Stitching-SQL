package StitchingSQLGo

import (
	"strconv"
)

type Limit uint64

func (l Limit) limit(s *sql) error {

	if s == nil {
		return ErrorNilSQL
	}
	if l == 0 {
		return nil
	}
	s.WriteString(" limit ")
	s.WriteString(strconv.FormatUint(uint64(l), 10))

	return nil
}
