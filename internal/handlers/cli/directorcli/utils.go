package directorcli

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

func PrintDirector(dir entities.Director) {
	fmt.Printf("[%v]: %s\n", dir.DirectorID, dir.Name)
}
