package sessioncli

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/services/sessionserv"
	"cinema/internal/repositories/sessionrepo"
	"cinema/internal/utils"
	"fmt"

	"gorm.io/gorm"
)

func Route(db *gorm.DB) {
	repo := sessionrepo.NewGormRepository(db)
	serv := sessionserv.New(repo)

out:
	for {
		selectedOption := utils.CrudMenuPrompt("Sessões")

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

func ListAll(service *sessionserv.Service) {
	var sessions []entities.Session
	var err error

out:
	for {
		selectedOption := SessionListingPrompt()

		switch selectedOption {
		case ListAllOption:
			sessions, err = service.ListAll()
			break out
		case ListByFilmOption:
			id := utils.StringPrompt("Digite o id do Filme:")
			sessions, err = service.GetByFilmId(id)
			break out
		case ListByBuyerOption:
			id := utils.StringPrompt("Digite o id do Pagante:")
			sessions, err = service.GetByBuyerId(id)
			break out
		case ListByGenderOption:
			id := utils.StringPrompt("Digite o id do Gênero:")
			sessions, err = service.GetByGenderId(id)
			break out
		case BackOption:
			return
		default:
			utils.PrintInvalidOption()
		}
	}

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if len(sessions) > 0 {
		fmt.Println("Sessões:")
		for _, session := range sessions {
			PrintSession(session)
		}
	} else {
		fmt.Println("Nenhum dado encontrado.")
	}
}

func Create(service *sessionserv.Service) {
	session := SessionCreationPrompt()
	err := service.Create(&session)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println("Sessão criado com sucesso.")
}

func Edit(service *sessionserv.Service) {
	id := utils.StringPrompt("Digite o id do Sessão:")
	session, err := service.Get(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	isSaveChanges := SessionEditPrompt(session)
	if !isSaveChanges {
		utils.PrintDiscardedChanges()
		return
	}

	updateErr := service.Update(session)
	if updateErr != nil {
		fmt.Printf("%v\n", updateErr)
		return
	}

	fmt.Println("Sessão atualizado com sucesso.")
}

func Delete(service *sessionserv.Service) {
	id := utils.StringPrompt("Digite o id do Sessão:")
	err := service.Delete(id)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println("Sessão excluído com sucesso.")
}
