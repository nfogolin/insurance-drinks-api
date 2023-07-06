package dto

type Aging struct {
	Aging *int32  `gorm:"column:aging" json:"aging"`
}
