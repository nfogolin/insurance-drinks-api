package dto

type BaseDrink struct {
	Id    string  `gorm:"column:drink_id;primary_key:true" json:"id" bson:"_id"`
	Name  string  `gorm:"column:name" json:"name" bson:"name" binding:"required"`
	Type  string  `gorm:"column:type" json:"type" bson:"type" binding:"required"`
	Price float64 `gorm:"column:price" json:"price" bson:"price" binding:"required"`
}
