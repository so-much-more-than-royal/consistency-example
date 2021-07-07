package servicec

type Storager interface {
	StoreA(i int)
	GetA() int
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
