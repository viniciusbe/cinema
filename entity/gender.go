package entity

import (
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	ID        uint
	Descricao string
}
