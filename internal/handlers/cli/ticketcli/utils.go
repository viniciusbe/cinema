package ticketcli

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/utils"
	"fmt"
	"regexp"
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
	fmt.Printf("ID: %v | Assento: %s | %s\n", ticket.TickedID, ticket.Seat, ticket.Modality)
	fmt.Printf("Pagante: %s\n", ticket.Buyer.Name)
	fmt.Printf("Id da sessão: %v\n", ticket.Session.SessionID)

	utils.PrintDivider()
}

func TicketCreationPrompt() entities.Ticket {
	ticket := entities.Ticket{}
	ticket.Seat = SeatPrompt()
	ticket.Modality = ModalityPrompt()
	ticket.Buyer.BuyerID = utils.IntPrompt("Id do pagante:")
	ticket.Session.SessionID = utils.IntPrompt("Id da sessão:")

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
			ticket.Seat = SeatPrompt()
		case EditModalityOption:
			ticket.Modality = ModalityPrompt()
		case EditBuyerIdOption:
			ticket.Buyer.BuyerID = utils.IntPrompt("Id do pagante:")
		case EditSessionIdOption:
			ticket.Session.SessionID = utils.IntPrompt("Id da sessão:")
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

func SeatPrompt() string {
	for {
		seat := utils.StringPrompt("Assento do ingresso (Ex: A16):")
		match, err := regexp.MatchString("^[A-P]([1-9]|1[0-8])$", seat)

		if match && err == nil {
			return seat
		} else {
			fmt.Println("Assento inválido, fileiras devem ser de A-P e cadeiras de 1-18.")
		}
	}
}
