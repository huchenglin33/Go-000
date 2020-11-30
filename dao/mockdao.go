package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

var ErrRecordNotFound = sql.ErrNoRows

type mockBiz struct {
}

func NewMockDao() *mockBiz {
	return &mockBiz{}
}

func (d *mockBiz) QueryCount() (int, error) {
	count := 0 //really count
	err := ErrRecordNotFound
	if isDBErr(err) {
		//db err
		return 0, errors.Wrap(err, "db query failed")
	}
	return count, nil
}

//TODO 补充err case
func isDBErr(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
