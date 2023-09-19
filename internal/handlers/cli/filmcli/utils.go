package filmcli

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/utils"
	"fmt"
	"strings"
)

const (
	ListAllOption = "1"
	CreateOption  = "2"
	EditOption    = "3"
	DeleteOption  = "4"
	BackOption    = "5"

	NoAgeFilm = 0

	MinGender       = 1
	GenderExitValue = 0
)

func PrintFilmLabel() {
	fmt.Println("[Id]: Nome | Duração(m) | Sinopse | Idade Class. | Nome Diretor")
}

func gendersToString(genders []entities.Gender) string {
	var s []string
	for _, gender := range genders {
		s = append(s, gender.Description)
	}

	return strings.Join(s, ", ")
}

func PrintFilm(film entities.Film) {
	genders := gendersToString(film.Genders)
	fmt.Printf("ID: %v | %s | %v minutos | %s\n", film.ID, film.Name, film.Duration, genders)
	fmt.Printf("Sinopse: %s\n", film.Synopsis)
	if film.Age == NoAgeFilm {
		fmt.Println("Classificação indicativa: Livre")
	} else {
		fmt.Printf("Classificação indicativa: %v anos\n", film.Age)
	}
	fmt.Printf("Diretor: %s\n", film.Director.Name)
	fmt.Println("----------------------------------------")
}

func FilmPrompt() (entities.Film, []uint) {
	film := entities.Film{}
	film.Name = utils.StringPrompt("Informe do nome do filme:")
	film.Duration = utils.IntPrompt("Informe a duração do filme em minutos:")
	film.Synopsis = utils.StringPrompt("Informe a sinopse do filme:")
	film.Age = utils.IntPrompt("Informe a idade indicativa do filme (0 para livre):")
	film.DirectorID = utils.IntPrompt("Informe o id do diretor do filme:")

	var gendersID []uint

	for {
		gender := utils.IntPrompt("Informe o id do gênero (0 para sair):")

		if gender == GenderExitValue {
			if len(gendersID) < MinGender {
				fmt.Println("Informe ao menos um gênero")
			} else {
				break
			}
		} else {
			gendersID = append(gendersID, gender)
		}
	}

	return film, gendersID
}
