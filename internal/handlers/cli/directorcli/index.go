package directorcli

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/services/directorserv"
	"cinema/internal/repositories/directorrepo"
	"cinema/internal/utils"
	"fmt"

	"gorm.io/gorm"
)

func Route(db *gorm.DB) {
	repo := directorrepo.NewGormRepository(db)
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
		fmt.Printf("Erro ao buscar Diretores -> %v\n", err)
		return
	}

	fmt.Println("[id]: Nome")
	for _, director := range directors {
		PrintDirector(director)
	}
}

func Create(service *directorserv.Service) {
	description := utils.StringPrompt("Digite o nome do Diretor:")
	director := entities.Director{Name: description}
	err := service.Create(&director)

	if err != nil {
		fmt.Printf("Erro ao buscar Diretor -> %v\n", err)
		return
	}
}

func Edit(service *directorserv.Service) {
	id := utils.StringPrompt("Digite o id do Diretor:")
	director, err := service.Get(id)

	if err != nil {
		fmt.Printf("Erro ao buscar Diretor -> %v\n", err)
		return
	}

	fmt.Println("[id]: Nome")
	PrintDirector(*director)

	description := utils.StringPrompt("Digite o nome do Diretor:")
	director.Name = description
	updateErr := service.Update(director)

	if updateErr != nil {
		fmt.Printf("Erro ao atualizar Diretor -> %v\n", updateErr)
		return
	}
	fmt.Println("Diretor atualizado com sucesso.")
}

func Delete(service *directorserv.Service) {
	id := utils.StringPrompt("Digite o id do Diretor:")
	err := service.Delete(id)

	if err != nil {
		fmt.Printf("Erro ao excluir Diretor -> %v\n", err)
		return
	}
	fmt.Println("Diretor excluído com sucesso.")
}
