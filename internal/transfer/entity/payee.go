package entity

type Payee struct {
	ID     string
	Amount uint
}

func NewPayee(ID string) Payee {
	return Payee{
		ID: ID,
	}
}
