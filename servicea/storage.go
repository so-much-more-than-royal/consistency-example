package servicea

import (
	"tmp/consistenter"
)

type storager interface {
	StoreA(i int)
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
