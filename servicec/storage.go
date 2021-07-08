package servicec

import "tmp/provider"

type storager interface {
	StoreA(i int)
	GetA() int
}

type ConsistentStorager interface {
	storager
	provider.Consistencer
}

type CommonStorager interface {
	storager
	WithConsistency() ConsistentStorager
}
