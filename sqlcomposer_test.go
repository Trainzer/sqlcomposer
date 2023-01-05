package sqlcomposer_test

import (
	"testing"

	"github.com/Trainzer/sqlcomposer"
	"github.com/stretchr/testify/assert"
)

func TestAddParam(t *testing.T) {

	t.Run("AddParam normal test", func(t *testing.T) {
		t.Parallel()

		id := 1
		gender := "M"

		sqlExpected := "select name from users where id = $1 and gender = $2"
		paramsExpected := []interface{}{id, gender}

		sqc := sqlcomposer.NewSqlComposer()
		sql := "select name from users where id = " + sqc.AddParam(id) + " and gender = " + sqc.AddParam(gender)

		assert.Equal(t, sqlExpected, sql)
		assert.Equal(t, paramsExpected, sqc.GetParams())
	})

	t.Run("AddParam NULL test", func(t *testing.T) {
		t.Parallel()

		id := 1
		email := "a@b.c"

		sqlExpected := "update users set email = $1, facebook = NULL where id = $2"
		paramsExpected := []interface{}{email, id}

		sqc := sqlcomposer.NewSqlComposer()
		sql := "update users set email = " + sqc.AddParam(email) + ", facebook = " + sqc.AddParam(nil) + " where id = " + sqc.AddParam(id)

		assert.Equal(t, sqlExpected, sql)
		assert.Equal(t, paramsExpected, sqc.GetParams())
	})
}
