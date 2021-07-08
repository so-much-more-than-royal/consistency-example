package consistenter

type Consistencer interface {
	Commit()
	Rollback()
}
