package filmcli

import (
	"cinema/internal/core/domain/entities"
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
			fmt.Println("Opção inválida.")
		}
	}
}

func ListAll(service *filmserv.Service) {
	films, err := service.ListAll()

	if err != nil {
		fmt.Printf("Erro ao buscar Filmes -> %v\n", err)
		return
	}

	fmt.Println("[id]: Nome")
	for _, film := range films {
		PrintFilm(film)
	}
}

func Create(service *filmserv.Service) {
	name := utils.StringPrompt("Digite o nome do filme:")
	film := entities.Film{Name: name}
	err := service.Create(&film)

	if err != nil {
		fmt.Printf("Erro ao buscar filme -> %v\n", err)
		return
	}
}

func Edit(service *filmserv.Service) {
	id := utils.StringPrompt("Digite o id do filme:")
	film, err := service.Get(id)

	if err != nil {
		fmt.Printf("Erro ao buscar filme -> %v\n", err)
		return
	}

	fmt.Println("[id]: Nome")
	PrintFilm(*film)

	description := utils.StringPrompt("Digite o nome do filme:")
	film.Name = description
	updateErr := service.Update(film)

	if updateErr != nil {
		fmt.Printf("Erro ao atualizar filme -> %v\n", updateErr)
		return
	}
	fmt.Println("filme atualizado com sucesso.")
}

func Delete(service *filmserv.Service) {
	id := utils.StringPrompt("Digite o id do filme:")
	err := service.Delete(id)

	if err != nil {
		fmt.Printf("Erro ao excluir filme -> %v\n", err)
		return
	}
	fmt.Println("filme excluído com sucesso.")
}
