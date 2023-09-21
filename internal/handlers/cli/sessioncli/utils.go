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

	timeLayout = "02/01/06 03:04"
)

func PrintSession(session entities.Session) {
	fmt.Printf("ID: %v | %s | %s\n", session.ID, session.Film.Name, session.Language)
	fmt.Printf("Sala: %v\n", session.Room)
	fmt.Printf("Data e hora (21/09/23 15:45): %s\n", session.Time.Local().Format(timeLayout))

	utils.PrintDivider()
}

func SessionPrompt() entities.Session {
	session := entities.Session{}
	session.FilmID = utils.IntPrompt("Id do filme:")
	language := utils.StringPrompt("Sessão (d)ublada ou (l)egendada:")
	if language == "d" {
		session.Language = "Dublado"
	} else {
		session.Language = "Legendado"
	}
	session.Room = utils.IntPrompt("Número da sala que irá passar:")
	session.Time = TimePrompt()

	return session
}

func TimePrompt() time.Time {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	value := utils.StringPrompt("Data e hora (21/09/23 15:45):\n")
	tm, _ := time.ParseInLocation(timeLayout, value, loc)
	return tm
}
