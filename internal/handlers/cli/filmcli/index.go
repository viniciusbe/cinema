package filmcli

import (
	"cinema/internal/core/services/filmserv"
	"cinema/internal/repositories/filmrepo"
	"cinema/internal/utils"
	"fmt"

	"gorm.io/gorm"
)

func Route(db *gorm.DB) {
	repo := filmrepo.NewGormRepository(db)
	serv := filmserv.New(repo)

out:
	for {
		selectedOption := utils.CrudMenuPrompt("Filmes")

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

func ListAll(service *filmserv.Service) {
	films, err := service.ListAll()

	if err != nil {
		fmt.Printf("Erro ao buscar Filmes -> %v\n", err)
		return
	}

	if len(films) > 0 {
		fmt.Println("Filmes:")
		for _, film := range films {
			PrintFilm(film)
		}
	} else {
		fmt.Println("Nenhum dado encontrado.")
	}
}

func Create(service *filmserv.Service) {
	film, gendersID := FilmPrompt()
	err := service.Create(&film, gendersID)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("Filme criado com sucesso.")
}

func Edit(service *filmserv.Service) {
	id := utils.StringPrompt("Digite o id do filme:")
	film, err := service.Get(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	gendersID, isSaveChanges := FilmEditPrompt(film)
	if !isSaveChanges {
		utils.PrintDiscardedChanges()
		return
	}

	updateErr := service.Update(film, gendersID)
	if updateErr != nil {
		fmt.Printf("%v\n", updateErr)
		return
	}

	fmt.Println("Filme atualizado com sucesso.")
}

func Delete(service *filmserv.Service) {
	id := utils.StringPrompt("Digite o id do filme:")
	film, err := service.Get(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	PrintFilm(*film)
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
	fmt.Println("Filme excluído com sucesso.")
}
