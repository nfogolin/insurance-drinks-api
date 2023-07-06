package entities

type Water struct {
	Drink
}

func (t *Water) GetTotalWithTaxes() float64 {
	return t.Price * 1.21
}

func (t *Water) GetAging() int32 {
	return 1
}
