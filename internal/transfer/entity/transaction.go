package entity

type Transaction struct {
	ID     string
	Payer  Payer
	Payee  Payee
	Amount uint
	Status string
}
