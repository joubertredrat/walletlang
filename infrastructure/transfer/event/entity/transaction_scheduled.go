package entity

type TransactionScheduled struct {
	Event string `json:"event"`
	ID    string `json:"id"`
}

func NewTransactionScheduled(ID string) TransactionScheduled {
	return TransactionScheduled{
		Event: "transaction_was_scheduled",
		ID:    ID,
	}
}
