package serviceb

import (
	"tmp/consistenter"
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
