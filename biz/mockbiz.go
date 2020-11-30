package biz

import (
	"github.com/pkg/errors"
	"jike/dao"
)

type mockBiz struct {
}

func NewMockBiz() *mockBiz {
	return &mockBiz{}
}

func (b *mockBiz) QueryCount() (int, error) {
	count, err := dao.NewMockDao().QueryCount()
	if err != nil {
		return 0, errors.WithMessagef(err, "model=biz||param=%s", "biz param")
	}
	return count, nil
}
