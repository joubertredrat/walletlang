package entity

const (
	TRANSACTION_STATUS_SCHEDULED       = "scheduled"
	TRANSACTION_STATUS_ERROR_SCHEDULED = "error_scheduled"
)

type Transaction struct {
	ID     string
	Payer  Payer
	Payee  Payee
	Amount uint
	Status string
}

func NewTransactionScheduled(Payer Payer, Payee Payee, Amount uint) Transaction {
	return Transaction{
		Payer:  Payer,
		Payee:  Payee,
		Amount: Amount,
		Status: TRANSACTION_STATUS_SCHEDULED,
	}
}

func (t *Transaction) SetStatusErrorScheduled() {
	t.Status = TRANSACTION_STATUS_ERROR_SCHEDULED
}
