package filmcli

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/utils"
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
	fmt.Printf("[%v]: %s | %v | %s | %v | %v\n", film.ID, film.Name, film.Duration, film.Synopsis, film.Age, film.DirectorID)
}

func FilmPrompt(film *entities.Film) {
	film.Name = utils.StringPrompt("Informe do nome do filme:")
	film.Duration = utils.IntPrompt("Informe a duração do filme:")
}
