package dao

type Transaction struct {
	ID      string `bson:"_id,omitempty"`
	AppID   string `bson:"app_id"`
	PayerID string `bson:"payer_id"`
	PayeeID string `bson:"payee_id"`
	Amount  uint   `bson:"amount"`
	Status  string `bson:"status"`
}
