package storage

// ConsistentStorage является реализацией domain.ConsistentStorager.
type ConsistentStorage struct {
}

func (c *ConsistentStorage) StoreA(i int) {
	panic("implement me")
}

func (c *ConsistentStorage) GetA() int {
	panic("implement me")
}

func (c *ConsistentStorage) StoreB(i int) {
	panic("implement me")
}

func (c *ConsistentStorage) GetB() int {
	panic("implement me")
}

func (c *ConsistentStorage) Commit() {
	panic("implement me")
}

func (c *ConsistentStorage) Rollback() {
	panic("implement me")
}
