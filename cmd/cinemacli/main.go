package main

import (
	"cinema/internal/handlers/cli/buyercli"
	"cinema/internal/handlers/cli/directorcli"
	"cinema/internal/handlers/cli/filmcli"
	"cinema/internal/handlers/cli/gendercli"
	"cinema/internal/handlers/neo4j"
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
	driver, ctx := neo4j.SetupNeo4jDB()

out:
	for {
		fmt.Printf("\nMenu do Cinema. Escolha uma opção:\n")
		fmt.Println("[1] Gêneros")
		fmt.Println("[2] Diretores")
		fmt.Println("[3] Filmes")
		// fmt.Println("[4] Sessões")
		fmt.Println("[5] Pagantes")
		// fmt.Println("[6] Ingressos")
		fmt.Printf("[s] Sair\n\n")

		selectedOption := utils.StringPrompt("")
		utils.PrintDivider()

		switch selectedOption {
		case GendersOption:
			gendercli.Route(driver, ctx)
		case DirectorsOption:
			directorcli.Route(driver, ctx)
		case FilmsOption:
			filmcli.Route(driver, ctx)
		// case SessionsOption:
		// 	sessioncli.Route(driver, ctx)
		case BuyersOption:
			buyercli.Route(driver, ctx)
		// case TicketsOption:
		// 	ticketcli.Route(driver, ctx)
		case ExitOption:
			break out
		default:
			utils.PrintInvalidOption()
		}
	}

	fmt.Println("Programa finalizado.")
}
