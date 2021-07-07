package serviceb

type Storager interface {
	StoreA(i int)
	GetA() int

	StoreB(i int)
	GetB(i int)
}

type ConsistentStorager interface {
	Storager
	Commit()
	Rollback()
}

type CommonStorager interface {
	Storager
	WithConsistency() ConsistentStorager
}
