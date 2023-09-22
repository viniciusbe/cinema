package ticketcli

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

	EditSeatOption      = "1"
	EditModalityOption  = "2"
	EditBuyerIdOption   = "3"
	EditSessionIdOption = "4"
	EditSaveOption      = "s"
	EditCancelOption    = "c"
	SaveChanges         = true
	DiscardChanges      = false
)

func PrintTicket(ticket entities.Ticket) {
	fmt.Printf("ID: %v | Assento: %s | %s\n", ticket.ID, ticket.Seat, ticket.Modality)
	fmt.Printf("Pagante: %s\n", ticket.Buyer.Name)
	fmt.Printf("Id da sessão: %v\n", ticket.SessionID)

	utils.PrintDivider()
}

func TicketCreationPrompt() entities.Ticket {
	ticket := entities.Ticket{}
	ticket.Seat = utils.StringPrompt("Assento do ingresso:")
	ticket.Modality = ModalityPrompt()
	ticket.BuyerID = utils.IntPrompt("Id do pagante:")
	ticket.SessionID = utils.IntPrompt("Id da sessão:")

	return ticket
}

func TicketEditPrompt(ticket *entities.Ticket) bool {
	PrintTicket(*ticket)
	for {
		fmt.Println("Edição de Ingresso, escolha uma opção:")
		fmt.Println("[1] Assento")
		fmt.Println("[2] Modalidade")
		fmt.Println("[3] Id do pagante")
		fmt.Println("[4] Id da sessão")
		fmt.Println("[s] Salvar")
		fmt.Println("[c] Cancelar")

		input := utils.StringPrompt("")

		switch input {
		case EditSeatOption:
			ticket.Seat = utils.StringPrompt("Assento do ingresso:")
		case EditModalityOption:
			ticket.Modality = ModalityPrompt()
		case EditBuyerIdOption:
			ticket.BuyerID = utils.IntPrompt("Id do pagante:")
		case EditSessionIdOption:
			ticket.SessionID = utils.IntPrompt("Id da sessão:")
		case EditSaveOption:
			return SaveChanges
		case EditCancelOption:
			return DiscardChanges
		default:
			utils.PrintInvalidOption()
		}
	}
}

func ModalityPrompt() string {
	language := utils.StringPrompt("Ingresso é (m)eia ou (i)nteira?")
	if language == "m" {
		return "Meia"
	} else {
		return "Inteira"
	}
}
