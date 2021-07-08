package storage

// consistentStorage является реализацией domain.ConsistentStorager.
type consistentStorage struct {
}

func (c *consistentStorage) StoreA(i int) {
	panic("implement me")
}

func (c *consistentStorage) GetA() int {
	panic("implement me")
}

func (c *consistentStorage) StoreB(i int) {
	panic("implement me")
}

func (c *consistentStorage) GetB() int {
	panic("implement me")
}

func (c *consistentStorage) Commit() {
	panic("implement me")
}

func (c *consistentStorage) Rollback() {
	panic("implement me")
}
