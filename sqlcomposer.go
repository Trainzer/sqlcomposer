package sqlcomposer

import "strconv"

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

func (sq *sqlcomposer) AddParam(in interface{}) string {
	if in == nil { //if in == nil || (reflect.ValueOf(in).Kind() == reflect.Ptr && reflect.ValueOf(in).IsNil())
		//reflect.Ptr(reflect.ValueOf(c)) && reflect.ValueOf(c).Elem().IsNil() ???
		return "NULL"
	}

	sq.paramList = append(sq.paramList, in)

	return sq.getNewParam()
}
