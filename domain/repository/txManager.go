package repository

type TransactionManagerI[Tx any] interface {
	Begin() Tx
	Commit(Tx) error
	Rollback(Tx) error
}
