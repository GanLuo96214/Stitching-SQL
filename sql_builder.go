package StitchingSQLGo

import (
	"strconv"
	"strings"
)

type SqlBuilder struct {
	strings.Builder
	args []interface{}
}

func (s *SqlBuilder) push(arg interface{}) error {
	if s == nil {
		return ErrorNilSQL
	}
	l := len(s.args) + 1
	lS := strconv.FormatInt(int64(l), 10)
	s.WriteString(" $")
	s.WriteString(lS)
	s.args = append(s.args, arg)
	return nil
}
