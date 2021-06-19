package entity

type Payer struct {
	ID     string
	Amount uint
}

func NewPayer(ID string) Payer {
	return Payer{
		ID: ID,
	}
}
