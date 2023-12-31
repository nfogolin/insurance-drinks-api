package entities

const(
	VINO = "Vino"
	WHISKY = "Whisky"
	WATER = "Water"
)

type Drink struct {
	Id    string  `gorm:"column:drink_id;primary_key:true" json:"id"`
	Name  string  `gorm:"column:name" json:"name" binding:"required"`
	Type  string  `gorm:"column:type" json:"type" binding:"required"`
	Price float64 `gorm:"column:price" json:"price" binding:"required"`
	Aging *int32   `gorm:"column:aging" json:"aging"`
}
