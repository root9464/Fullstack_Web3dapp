package model

import "github.com/lib/pq"

type Record struct {
	Name    string        `gorm:"primaryKey;not null;unique"`
	Total   pq.Int64Array `gorm:"type:int[];not null"`
	Percent int64         `gorm:"not null"`
}
