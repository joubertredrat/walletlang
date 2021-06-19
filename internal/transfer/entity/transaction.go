package entity

const (
	TRANSACTION_STATUS_SCHEDULED          = "scheduled"
	TRANSACTION_STATUS_ERROR_SCHEDULED    = "error_scheduled"
	TRANSACTION_STATUS_INSUFFICIENT_FUNDS = "insufficient_funds"
	TRANSACTION_STATUS_DONE               = "done"
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

func (t *Transaction) SetStatusInsufficientFunds() {
	t.Status = TRANSACTION_STATUS_INSUFFICIENT_FUNDS
}

func (t *Transaction) SetStatusDone() {
	t.Status = TRANSACTION_STATUS_DONE
}
