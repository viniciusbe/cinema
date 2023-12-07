package buyercli

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/services/buyerserv"
	"cinema/internal/repositories/neo4jdb/buyerrepo"
	"cinema/internal/utils"
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func Route(driver neo4j.DriverWithContext, ctx context.Context) {
	repo := buyerrepo.NewNeo4jRepository(driver, ctx)
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

	if len(buyers) > 0 {
		fmt.Println("Pagantes:")
		PrintBuyerLabel()
		for _, buyer := range buyers {
			PrintBuyer(buyer)
		}
	} else {
		fmt.Println("Nenhum dado encontrado.")
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
	buyer, err := service.Get(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	PrintBuyerLabel()
	PrintBuyer(*buyer)
	confirmDeletion := utils.ConfirmDeletePrompt()
	if !confirmDeletion {
		fmt.Println("Exclusão cancelada.")
		return
	}

	err = service.Delete(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("Pagante excluido com sucesso.")
}
