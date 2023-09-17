package main

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/handlers"
	gormdb "cinema/internal/handlers"
	"cinema/internal/utils"
	"fmt"
)

const (
	GendersOption   = "1"
	DirectorsOption = "2"
	FilmsOption     = "3"
	SessionsOption  = "4"
	TicketsOption   = "5"
	BuyersOption    = "6"
	ExitOption      = "s"
)

func main() {
	db, err := gormdb.SetupGormDB()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.Gender{})

out:
	for {
		fmt.Printf("\nMenu do Cinema. Escolha uma opção:\n")
		fmt.Println("  [1] Genders")
		fmt.Println("  [2] Directors")
		fmt.Println("  [3] Films")
		fmt.Println("  [4] Sessions")
		fmt.Println("  [5] Buyers")
		fmt.Println("  [6] Tickets")
		fmt.Printf("  [s] Sair\n\n")

		selectedOption := utils.StringPrompt()

		switch selectedOption {
		case GendersOption:
			handlers.Gender(db)
		case DirectorsOption:
		case FilmsOption:
		case SessionsOption:
		case BuyersOption:
		case TicketsOption:
		case ExitOption:
			break out
		default:
			fmt.Println("Opção inválida.")
		}
	}

	fmt.Println("Programa finalizado.")
}
