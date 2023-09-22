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

	EditNameOption     = "1"
	EditDurationOption = "2"
	EditSynopsisOption = "3"
	EditAgeOption      = "4"
	EditDirectorOption = "5"
	EditGendersOption  = "6"
	EditSaveOption     = "s"
	EditCancelOption   = "c"
	SaveChanges        = true
	DiscardChanges     = false

	AddGenderOption    = "1"
	RemoveGenderOption = "2"

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
	utils.PrintDivider()
}

func FilmPrompt() (entities.Film, []uint) {
	film := entities.Film{}
	film.Name = utils.StringPrompt("Nome do filme:")
	film.Duration = utils.IntPrompt("Duração do filme em minutos:")
	film.Synopsis = utils.StringPrompt("Sinopse do filme:")
	film.Age = utils.IntPrompt("Idade da class. indicativa do filme (0 para livre):")
	film.DirectorID = utils.IntPrompt("Id do diretor do filme:")
	gendersID := GenderPrompt("Informe o id do gênero (0 para voltar):", true)

	return film, gendersID
}

func FilmEditPrompt(film *entities.Film) ([]uint, bool) {
	PrintFilm(*film)

	var gendersIDToUpdate []uint

	for {
		fmt.Println("Edição de Filme, escolha uma opção:")
		fmt.Println("[1] Nome")
		fmt.Println("[2] Duração")
		fmt.Println("[3] Sinopse")
		fmt.Println("[4] Classificação indicativa")
		fmt.Println("[5] Id do diretor")
		fmt.Println("[6] Gêneros")
		fmt.Println("[s] Salvar")
		fmt.Println("[c] Cancelar")

		input := utils.StringPrompt("")

		switch input {
		case EditNameOption:
			film.Name = utils.StringPrompt("Nome do filme:")
		case EditDurationOption:
			film.Duration = utils.IntPrompt("Duração do filme em minutos:")
		case EditSynopsisOption:
			film.Synopsis = utils.StringPrompt("Sinopse do filme:")
		case EditAgeOption:
			film.Age = utils.IntPrompt("Idade da class. indicativa do filme (0 para livre):")
		case EditDirectorOption:
			film.DirectorID = utils.IntPrompt("ID do diretor do filme:")
		case EditGendersOption:
			fmt.Println("Informe os ids dos novos gêneros (todos os gêneros atuais serão sobrescritos):")
			gendersIDToUpdate = GenderPrompt("ID do gênero (0 para sair):", false)
		case EditSaveOption:
			return gendersIDToUpdate, SaveChanges
		case EditCancelOption:
			fmt.Println("Alterações descartadas")
			return nil, DiscardChanges
		default:
			utils.PrintInvalidOption()
		}
	}
}

func GenderPrompt(label string, isCreate bool) []uint {
	var gendersID []uint
	for {
		gender := utils.IntPrompt(label)
		gendersID = append(gendersID, gender)

		if gender == GenderExitValue {
			hasNotMinGender := len(gendersID) < MinGender

			if hasNotMinGender && isCreate {
				fmt.Println("Informe ao menos um gênero.")
			} else {
				break
			}
		}
	}

	return gendersID
}
