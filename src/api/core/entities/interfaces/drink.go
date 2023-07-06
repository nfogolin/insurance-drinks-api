package interfaces

type Drink interface {
	GetTotalWithTaxes() float64
	GetAging() int32
}
