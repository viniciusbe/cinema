package filmcli

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

func PrintFilm(film entities.Film) {
	fmt.Printf("[%v]: %s\n", film.ID, film.Name)
}
