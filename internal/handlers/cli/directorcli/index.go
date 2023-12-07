package directorcli

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/services/directorserv"
	"cinema/internal/repositories/neo4jdb/directorrepo"
	"cinema/internal/utils"
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func Route(driver neo4j.DriverWithContext, ctx context.Context) {
	repo := directorrepo.NewNeo4jRepository(driver, ctx)
	serv := directorserv.New(repo)

out:
	for {
		selectedOption := utils.CrudMenuPrompt("Diretores")

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
			fmt.Println("Opção inválida.")
		}
	}
}

func ListAll(service *directorserv.Service) {
	directors, err := service.ListAll()

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if len(directors) > 0 {
		fmt.Println("Diretores:")
		fmt.Println("[id]: Nome")
		for _, director := range directors {
			PrintDirector(director)
		}
	} else {
		fmt.Println("Nenhum dado encontrado.")
	}
}

func Create(service *directorserv.Service) {
	description := utils.StringPrompt("Digite o nome do Diretor:")
	director := entities.Director{Name: description}
	err := service.Create(&director)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println("Diretor criado com sucesso.")
}

func Edit(service *directorserv.Service) {
	id := utils.StringPrompt("Digite o id do Diretor:")
	director, err := service.Get(id)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("[id]: Nome")
	PrintDirector(*director)

	description := utils.StringPrompt("Digite o nome do Diretor:")
	director.Name = description
	updateErr := service.Update(director)

	if updateErr != nil {
		fmt.Printf("%v\n", updateErr)
		return
	}
	fmt.Println("Diretor atualizado com sucesso.")
}

func Delete(service *directorserv.Service) {
	id := utils.StringPrompt("Digite o id do Diretor:")
	director, err := service.Get(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("[id]: Nome")
	PrintDirector(*director)

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
	fmt.Println("Diretor excluído com sucesso.")
}
