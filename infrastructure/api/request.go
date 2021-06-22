package api

type TransactionScheduleRequest struct {
	PayerID string `json:"payer_id"`
	PayeeID string `json:"payee_id"`
	Amount  uint   `json:"amount"`
}
