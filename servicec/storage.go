package servicec

import "tmp/domain"

type storager interface {
	StoreA(i int)
	GetA() int
}

type ConsistentStorager interface {
	storager
	domain.Consistencer
}

type CommonStorager interface {
	storager
	WithConsistency() ConsistentStorager
}
