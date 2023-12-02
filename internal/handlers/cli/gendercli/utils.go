package gendercli

import (
	"cinema/internal/core/domain/entities"
	"fmt"
)

const (
	ListAllOption = "1"
	CreateOption  = "2"
	EditOption    = "3"
	DeleteOption  = "4"
	BackOption    = "5"
)

func PrintGender(gender entities.Gender) {
	fmt.Printf("[%v]: %s\n", gender.GenderID, gender.Description)
}
