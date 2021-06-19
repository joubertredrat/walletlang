package entity

type Payer struct {
	ID string
}

func NewPayer(ID string) Payer {
	return Payer{
		ID: ID,
	}
}
