package servicea

import (
	"tmp/provider"
	"tmp/storage"
)

type storager interface {
	StoreA(i int)
	GetB() int
}

type ConsistentStorager interface {
	storager
	provider.Consistencer
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
