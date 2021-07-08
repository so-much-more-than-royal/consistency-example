package storage

// commonStorage является реализацией интерфейса domain.CommonStorager.
type commonStorage struct {
}

func (c *commonStorage) StoreA(i int) {
	panic("implement me")
}

func (c *commonStorage) GetA() int {
	panic("implement me")
}

func (c *commonStorage) StoreB(i int) {
	panic("implement me")
}

func (c *commonStorage) GetB() int {
	panic("implement me")
}

func (c *commonStorage) WithConsistency() *consistentStorage {
	panic("implement me")
}
