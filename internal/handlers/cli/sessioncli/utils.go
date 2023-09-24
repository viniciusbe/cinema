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

	MinValidRoom = 1
	MaxVadidRoom = 7
)

func PrintSession(session entities.Session) {
	fmt.Printf("ID: %v | %s | %s\n", session.ID, session.Film.Name, session.Language)
	fmt.Printf("Sala: %v\n", session.Room)
	fmt.Printf("Data e hora: %s\n", FormatTime(session.Time))

	utils.PrintDivider()
}

func SessionCreationPrompt() entities.Session {
	session := entities.Session{}
	session.FilmID = utils.IntPrompt("Id do filme:")
	session.Language = LanguagePrompt()
	session.Room = RoomPrompt()
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
			session.Room = RoomPrompt()
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
	for {
		fmt.Printf("Informe a data e hora (Ex: %s):\n", FormatTime(time.Now()))
		value := utils.StringPrompt("")
		tm, err := time.ParseInLocation(timeLayout, value, loc)

		startTime := time.Now()
		endTime := time.Now().AddDate(0, 1, 0)
		isValidDate := tm.After(startTime) && tm.Before(endTime)

		if err == nil && isValidDate {
			return tm
		} else {
			fmt.Printf("Data inválida. A data deve ser entre %s e %s.\n", FormatTime(startTime), FormatTime(endTime))
		}
	}
}

func LanguagePrompt() string {
	language := utils.StringPrompt("Sessão (d)ublada ou (l)egendada?")
	if language == "d" {
		return "Dublado"
	} else {
		return "Legendado"
	}
}

func RoomPrompt() uint {
	for {
		fmt.Printf("Número da sala (de %v a %v):\n", MinValidRoom, MaxVadidRoom)
		room := utils.IntPrompt("")
		if room < MinValidRoom || room > MaxVadidRoom {
			fmt.Println("Sala não existe.")
		} else {
			return room
		}
	}

}

func FormatTime(time time.Time) string {
	return time.Local().Format(timeLayout)
}
