package storage

import (
	"tmp/domain"
)

// commonStorage является реализацией интерфейса domain.CommonStorager.
type commonStorage struct {
}

func (c *commonStorage) StoreA(i int) {
	panic("implement me")
}

func (c *commonStorage) GetA() int {
	panic("implement me")
}

func (c *commonStorage) StoreB(i int) {
	panic("implement me")
}

func (c *commonStorage) GetB(i int) {
	panic("implement me")
}

func (c *commonStorage) WithConsistency() domain.ConsistentStorager {
	panic("implement me")
}
