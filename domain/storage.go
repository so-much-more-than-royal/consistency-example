package domain

type Consistencer interface {
	Commit()
	Rollback()
}
