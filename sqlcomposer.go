package sqlcomposer

import (
	"reflect"
	"strconv"
	"strings"
)

const (
	SQLNullValue = "NULL"
)

type sqlcomposer struct {
	paramCounter int
	paramList    []interface{}
}

func NewSqlComposer() *sqlcomposer {
	return &sqlcomposer{paramCounter: 0}
}

func (sq *sqlcomposer) GetParams() []interface{} {
	return sq.paramList
}

func (sq *sqlcomposer) getNewParam() string {
	sq.paramCounter++
	return "$" + strconv.Itoa(sq.paramCounter)
}

// Add parameter to query
func (sq *sqlcomposer) AddParam(in interface{}) string {
	if in == nil {
		return SQLNullValue
	}

	sq.paramList = append(sq.paramList, in)

	return sq.getNewParam()
}

// If isNull, write NULL value
func (sq *sqlcomposer) AddNullableParam(in interface{}, isNull bool) string {
	if isNull {
		return SQLNullValue
	}

	sq.paramList = append(sq.paramList, in)

	return sq.getNewParam()
}

// Expand array to variable list. See test for package
func (sq *sqlcomposer) AddArrayParam(in interface{}) string {
	s := reflect.ValueOf(in)
	switch s.Kind() {
	case reflect.Slice, reflect.Array:
		if s.Len() == 0 {
			return SQLNullValue
		}

		s := reflect.ValueOf(in)
		sb := strings.Builder{}

		for i := 0; i < s.Len(); i++ {
			if i > 0 {
				sb.WriteRune(rune(','))
			}
			sb.WriteString(sq.AddParam(s.Index(i).Interface()))
		}
		return sb.String()
	default:
		return SQLNullValue
	}
}

func (sq *sqlcomposer) Ife(condition bool, ifTrue, ifFalse string) string {
	if condition {
		return ifTrue
	}
	return ifFalse
}

func (sq *sqlcomposer) If(condition bool, ifTrue string) string {
	if condition {
		return ifTrue
	}
	return ""
}

func (sq *sqlcomposer) IfeF(condition bool, ifTrue, ifFalse func() string) string {
	if condition {
		return ifTrue()
	}
	return ifFalse()
}

func (sq *sqlcomposer) IfF(condition bool, ifTrue func() string) string {
	if condition {
		return ifTrue()
	}
	return ""
}
