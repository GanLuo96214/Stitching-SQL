package StitchingSQLGo

import (
	"strconv"
)

type Offset uint64

func (o Offset) offset(s *sql) error {

	if o == 0 {
		return nil
	}
	s.WriteString(" offset ")
	s.WriteString(strconv.FormatUint(uint64(o), 10))

	return nil
}
