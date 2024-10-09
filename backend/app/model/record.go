package model

type Record struct {
	Name    string  `gorm:"primaryKey;not null;unique"`
	Total   float64 `gorm:"not null"`
	Percent float64 `gorm:"not null"`
}
