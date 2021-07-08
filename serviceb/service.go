package serviceb

type service struct {
	storage CommonStorager
}

func New(storage CommonStorager) *service {
	return &service{
		storage: storage,
	}
}

func (s *service) DoSome(i int) {
	// Работаем в обычном режиме
	a := s.storage.GetA()
	b := s.storage.GetB()

	// Работаем в транзакционном режиме
	tx := s.storage.WithConsistency()
	defer tx.Rollback()

	a = tx.GetA()
	b = tx.GetB()
	tx.StoreA(a + 1)
	tx.StoreB(b + 1)

	tx.Commit()
}
