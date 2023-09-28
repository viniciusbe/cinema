package gendercli

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/services/genderserv"
	"cinema/internal/repositories/genderrepo"
	"cinema/internal/utils"
	"fmt"

	"gorm.io/gorm"
)

func Route(db *gorm.DB) {
	repo := genderrepo.NewGormRepository(db)
	serv := genderserv.New(repo)

out:
	for {
		selectedOption := utils.CrudMenuPrompt("Gêneros")

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

func ListAll(service *genderserv.Service) {
	genders, err := service.ListAll()

	if err != nil {
		fmt.Printf("Erro ao buscar Gêneros -> %v\n", err)
		return
	}

	if len(genders) > 0 {
		fmt.Println("Gêneros:")
		fmt.Println("[id]: Descrição")
		for _, gender := range genders {
			PrintGender(gender)
		}
	} else {
		fmt.Println("Nenhum dado encontrado.")
	}
}

func Create(service *genderserv.Service) {
	description := utils.StringPrompt("Digite o nome do Gênero:")
	gender := entities.Gender{Description: description}
	err := service.Create(&gender)

	if err != nil {
		fmt.Printf("Erro ao criar Gênero -> %v\n", err)
		return
	}
	fmt.Println("Gênero criado com sucesso.")
}

func Edit(service *genderserv.Service) {
	id := utils.StringPrompt("Digite o id do Gênero:")
	gender, err := service.Get(id)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("[id]: Descrição")
	PrintGender(*gender)

	description := utils.StringPrompt("Digite o nome do Gênero:")
	gender.Description = description
	updateErr := service.Update(gender)

	if updateErr != nil {
		fmt.Printf("Erro ao atualizar Gênero -> %v\n", updateErr)
		return
	}
	fmt.Println("Gênero atualizado com sucesso.")
}

func Delete(service *genderserv.Service) {
	id := utils.StringPrompt("Digite o id do Genero:")
	gender, err := service.Get(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("[id]: Descrição")
	PrintGender(*gender)

	confirmDeletion := utils.ConfirmDeletePrompt()
	if !confirmDeletion {
		fmt.Println("Exclusão cancelada.")
		return
	}

	err = service.Delete(id)
	if err != nil {
		fmt.Printf("Erro ao excluir Gênero -> %v\n", err)
		return
	}
	fmt.Println("Gênero excluido com sucesso.")
}
