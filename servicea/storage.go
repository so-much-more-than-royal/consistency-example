package servicea

import "tmp/provider"

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
