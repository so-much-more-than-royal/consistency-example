package storage

// CommonStorage является реализацией интерфейса domain.CommonStorager.
type CommonStorage struct {
}

func (c *CommonStorage) StoreA(i int) {
	panic("implement me")
}

func (c *CommonStorage) GetA() int {
	panic("implement me")
}

func (c *CommonStorage) StoreB(i int) {
	panic("implement me")
}

func (c *CommonStorage) GetB() int {
	panic("implement me")
}

func (c *CommonStorage) WithConsistency() *ConsistentStorage {
	panic("implement me")
}
