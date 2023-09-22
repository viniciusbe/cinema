package sessioncli

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/utils"
	"fmt"
	"time"
)

const (
	ListAllOption = "1"
	CreateOption  = "2"
	EditOption    = "3"
	DeleteOption  = "4"
	BackOption    = "5"

	timeLayout = "02/01/06 15:04"

	EditFilmIdOption   = "1"
	EditLanguageOption = "2"
	EditRoomOption     = "3"
	EditTimeOption     = "4"
	EditSaveOption     = "s"
	EditCancelOption   = "c"
	SaveChanges        = true
	DiscardChanges     = false
)

func PrintSession(session entities.Session) {
	fmt.Printf("ID: %v | %s | %s\n", session.ID, session.Film.Name, session.Language)
	fmt.Printf("Sala: %v\n", session.Room)
	fmt.Printf("Data e hora: %s\n", session.Time.Local().Format(timeLayout))

	utils.PrintDivider()
}

func SessionCreationPrompt() entities.Session {
	session := entities.Session{}
	session.FilmID = utils.IntPrompt("Id do filme:")
	session.Language = LanguagePrompt()
	session.Room = utils.IntPrompt("Número da sala que irá passar:")
	session.Time = TimePrompt()

	return session
}

func SessionEditPrompt(session *entities.Session) bool {
	PrintSession(*session)
	for {
		fmt.Println("Edição de Sessão, escolha uma opção:")
		fmt.Println("[1] Id do filme")
		fmt.Println("[2] Linguagem")
		fmt.Println("[3] Número da sala")
		fmt.Println("[4] Data e hora ")
		fmt.Println("[s] Salvar")
		fmt.Println("[c] Cancelar")

		input := utils.StringPrompt("")

		switch input {
		case EditFilmIdOption:
			session.FilmID = utils.IntPrompt("Id do filme:")
		case EditLanguageOption:
			session.Language = LanguagePrompt()
		case EditRoomOption:
			session.Room = utils.IntPrompt("Número da sala que irá passar:")
		case EditTimeOption:
			session.Time = TimePrompt()
		case EditSaveOption:
			return true
		case EditCancelOption:
			return false
		default:
			utils.PrintInvalidOption()
		}
	}
}

func TimePrompt() time.Time {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	value := utils.StringPrompt("Data e hora (21/09/23 15:45):")
	tm, _ := time.ParseInLocation(timeLayout, value, loc)
	return tm
}

func LanguagePrompt() string {
	language := utils.StringPrompt("Sessão (d)ublada ou (l)egendada?")
	if language == "d" {
		return "Dublado"
	} else {
		return "Legendado"
	}
}