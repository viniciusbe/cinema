package main

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/handlers/cli/buyercli"
	"cinema/internal/handlers/cli/directorcli"
	"cinema/internal/handlers/cli/filmcli"
	"cinema/internal/handlers/cli/gendercli"
	"cinema/internal/handlers/cli/sessioncli"
	"cinema/internal/handlers/cli/ticketcli"
	"cinema/internal/handlers/gormdb"
	"cinema/internal/utils"
	"fmt"
)

const (
	GendersOption   = "1"
	DirectorsOption = "2"
	FilmsOption     = "3"
	SessionsOption  = "4"
	BuyersOption    = "5"
	TicketsOption   = "6"
	ExitOption      = "s"
)

func main() {
	db, err := gormdb.SetupGormDB()

	if err != nil {
		fmt.Printf("Erro na conexão com o banco -> %v\n", err)
		panic(err)
	}

	db.AutoMigrate(&entities.Gender{}, &entities.Director{}, &entities.Film{}, &entities.Session{}, &entities.Buyer{}, &entities.Ticket{})

out:
	for {
		fmt.Printf("\nMenu do Cinema. Escolha uma opção:\n")
		fmt.Println("[1] Gêneros")
		fmt.Println("[2] Diretores")
		fmt.Println("[3] Filmes")
		fmt.Println("[4] Sessões")
		fmt.Println("[5] Pagantes")
		fmt.Println("[6] Ingressos")
		fmt.Printf("[s] Sair\n\n")

		selectedOption := utils.StringPrompt("")
		utils.PrintDivider()

		switch selectedOption {
		case GendersOption:
			gendercli.Route(db)
		case DirectorsOption:
			directorcli.Route(db)
		case FilmsOption:
			filmcli.Route(db)
		case SessionsOption:
			sessioncli.Route(db)
		case BuyersOption:
			buyercli.Route(db)
		case TicketsOption:
			ticketcli.Route(db)
		case ExitOption:
			break out
		default:
			utils.PrintInvalidOption()
		}
	}

	fmt.Println("Programa finalizado.")
}
