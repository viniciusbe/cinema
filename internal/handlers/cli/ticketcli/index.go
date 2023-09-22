package ticketcli

import (
	"cinema/internal/core/services/ticketserv"
	"cinema/internal/repositories/ticketrepo"
	"cinema/internal/utils"
	"fmt"

	"gorm.io/gorm"
)

func Route(db *gorm.DB) {
	repo := ticketrepo.NewGormRepository(db)
	serv := ticketserv.New(repo)

out:
	for {
		selectedOption := utils.CrudMenuPrompt("Ingressos")

		switch selectedOption {
		case ListAllOption:
			ListAll(serv)
		case CreateOption:
			Create(serv)
		case EditOption:
			Edit(serv)
		case DeleteOption:
			Delete(serv)
		case BackOption:
			break out
		default:
			utils.PrintInvalidOption()
		}
	}
}

func ListAll(service *ticketserv.Service) {
	tickets, err := service.ListAll()

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("Ingressos:")
	for _, ticket := range tickets {
		PrintTicket(ticket)
	}
}

func Create(service *ticketserv.Service) {
	ticket := TicketCreationPrompt()
	err := service.Create(&ticket)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println("Ingresso criado com sucesso.")
}

func Edit(service *ticketserv.Service) {
	id := utils.StringPrompt("Digite o id do Ingresso:")
	ticket, err := service.Get(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	isSaveChanges := TicketEditPrompt(ticket)
	if !isSaveChanges {
		utils.PrintDiscardedChanges()
		return
	}

	updateErr := service.Update(ticket)
	if updateErr != nil {
		fmt.Printf("%v\n", updateErr)
		return
	}

	fmt.Println("Ingresso atualizado com sucesso.")
}

func Delete(service *ticketserv.Service) {
	id := utils.StringPrompt("Digite o id do Ingresso:")
	err := service.Delete(id)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println("Ingresso excluído com sucesso.")
}
