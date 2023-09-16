package domain

import (
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	ID        uint
	Descricao string
}
