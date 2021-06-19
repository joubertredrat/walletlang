package entity

type Payee struct {
	ID string
}

func NewPayee(ID string) Payee {
	return Payee{
		ID: ID,
	}
}
