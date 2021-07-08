package serviceb

import "tmp/domain"

type storager interface {
	StoreA(i int)
	GetA() int

	StoreB(i int)
	GetB() int
}

type ConsistentStorager interface {
	storager
	domain.Consistencer
}

type CommonStorager interface {
	storager
	WithConsistency() ConsistentStorager
}
