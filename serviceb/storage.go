package serviceb

import (
	"tmp/consistenter"
	"tmp/storage"
)

type storager interface {
	StoreA(i int)
	GetA() int

	StoreB(i int)
	GetB() int
}

type ConsistentStorager interface {
	storager
	consistenter.Consistencer
}

type CommonStorager interface {
	storager
	WithConsistency() ConsistentStorager
}

func NewCommonStorage(storage *storage.CommonStorage) CommonStorager {
	return &commonStorage{storage}
}

type commonStorage struct {
	*storage.CommonStorage
}

func (c *commonStorage) WithConsistency() ConsistentStorager {
	return c.CommonStorage.WithConsistency()
}
