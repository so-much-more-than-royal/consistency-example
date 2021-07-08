package servicec

import (
	"tmp/consistenter"
)

type storager interface {
	StoreA(i int)
	GetA() int
}

type ConsistentStorager interface {
	storager
	consistenter.Consistencer
}

type CommonStorager interface {
	storager
	WithConsistency() ConsistentStorager
}
