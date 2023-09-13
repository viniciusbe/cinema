package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StringPrompt() string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func main() {
	input := ""

	for {
		fmt.Println("Cinema - Menu principal:")

		fmt.Println("[1] Generos")
		fmt.Println("[1] Filmes")
		fmt.Println("[1] Diretores")
		fmt.Println("[1] Ingressos")

		input = StringPrompt()

		if input == "s" {
			fmt.Println("Programa finalizado")
			break
		}

		for {
			fmt.Println("[1] Listar todos")
			fmt.Println("[2] Criar")
			fmt.Println("[3] Editar")
			fmt.Println("[4] Excluir")
			fmt.Println("[5] Voltar")

			input = StringPrompt()

			if input == "5" {
				break
			}

			switch input {
			case "1":
				fmt.Println("Listar")
			case "2":
				fmt.Println("Criar")
			case "3":
				fmt.Println("Editar")
			case "4":
				fmt.Println("Excluir")
			}

		}
	}

	// descricaoGenero := StringPrompt()

	// db, err := gormdb.SetupGormDB()

	// if err != nil {
	// 	panic(err)
	// }

	// db.AutoMigrate(&entity.Gender{})

	// newGender := entity.Gender{Descricao: descricaoGenero}

	// db.Create(&newGender)

	// var matchingGender entity.Gender
	// // var genders []entity.Gender
	// db.First(&matchingGender, newGender.ID)

	// fmt.Println(matchingGender.Descricao)

}
