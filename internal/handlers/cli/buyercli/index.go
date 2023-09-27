package buyercli

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/services/buyerserv"
	"cinema/internal/repositories/buyerrepo"
	"cinema/internal/utils"
	"fmt"

	"gorm.io/gorm"
)

func Route(db *gorm.DB) {
	repo := buyerrepo.NewGormRepository(db)
	serv := buyerserv.New(repo)

out:
	for {
		selectedOption := utils.CrudMenuPrompt("Pagantes")

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

func ListAll(service *buyerserv.Service) {
	buyers, err := service.ListAll()

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("Pagantes:")
	PrintBuyerLabel()
	for _, buyer := range buyers {
		PrintBuyer(buyer)
	}
}

func Create(service *buyerserv.Service) {
	name := utils.StringPrompt("Nome do Pagante:")
	document := utils.StringPrompt("Documento do Pagante:")
	buyer := entities.Buyer{Name: name, Document: document}
	err := service.Create(&buyer)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("Pagante criado com sucesso.")
}

func Edit(service *buyerserv.Service) {
	id := utils.StringPrompt("Id do Pagante:")
	buyer, err := service.Get(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	isSaveChanges := BuyerEditPrompt(buyer)
	if !isSaveChanges {
		utils.PrintDiscardedChanges()
		return
	}

	updateErr := service.Update(buyer)
	if updateErr != nil {
		fmt.Printf("%v\n", updateErr)
		return
	}

	fmt.Println("Pagante atualizado com sucesso.")
}

func Delete(service *buyerserv.Service) {
	id := utils.StringPrompt("Id do Pagante:")
	err := service.Delete(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("Pagante excluido com sucesso.")
}
