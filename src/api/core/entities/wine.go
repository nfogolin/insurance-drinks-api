package entities

type Wine struct {
	Drink
}

func (t *Wine) GetTotalWithTaxes() float64 {
	return t.Price * 1.27
}

func (t *Wine) GetAging() int32 {
	return 1
}
