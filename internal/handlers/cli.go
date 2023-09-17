package handlers

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/services"
	"cinema/internal/repositories"
	"cinema/internal/utils"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

const (
	ListAllOption = "1"
	CreateOption  = "2"
	EditOption    = "3"
	DeleteOption  = "4"
	BackOption    = "5"
)

func CrudMenuPrompt(label string) string {
	fmt.Printf("\n  Menu de %s. Escolha uma opção:\n", label)
	fmt.Println("    [1] Listar todos")
	fmt.Println("    [2] Criar")
	fmt.Println("    [3] Editar")
	fmt.Println("    [4] Excluir")
	fmt.Printf("    [5] Voltar\n\n")

	input := utils.StringPrompt()

	return input
}

func Gender(db *gorm.DB) {
	genderRepo := repositories.NewGormGenderRepository(db)
	genderService := services.NewGenderService(genderRepo)

out:
	for {
		selectedOption := CrudMenuPrompt("Gêneros")

		switch selectedOption {
		case ListAllOption:
			genders, err := genderService.ListAll()

			if err != nil {
				break
			}

			for _, gender := range genders {
				fmt.Printf("[%v]: %s\n", gender.ID, gender.Descricao)
			}

		case CreateOption:
			fmt.Println("Digite o nome do Gênero:")
			description := utils.StringPrompt()
			gender := entities.Gender{Descricao: description}
			genderService.Create(&gender)

		case EditOption:
			fmt.Println("Digite o id do Genero:")
			id := utils.StringPrompt()

			fmt.Println("Digite o nome do Gênero:")
			description := utils.StringPrompt()
			i, _ := strconv.Atoi(id)
			gender := entities.Gender{ID: uint(i), Descricao: description}
			genderService.Update(&gender)

		case DeleteOption:
			fmt.Println("Digite o id do Genero:")
			id := utils.StringPrompt()

			genderService.Delete(id)
		case BackOption:
			break out
		default:
			fmt.Println("Opção inválida.")
		}
	}

}
