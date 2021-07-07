package domain

// Storager предоставляет функционал хранилища сущностей.
type Storager interface {
	StoreA(i int)
	GetA() int

	StoreB(i int)
	GetB(i int)
}

// ConsistentStorager предоставляет функционал консистентного хранилища сущностей.
type ConsistentStorager interface {
	Storager
	Consistencer
}

// CommonStorager предоставляет функционал хранилища сущностей, из которого можно получить консистентное хранилище.
type CommonStorager interface {
	Storager
	WithConsistency() ConsistentStorager
}

type Consistencer interface {
	Commit()
	Rollback()
}
