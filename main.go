package main

import (
	"bufio"
	"cinema/entity"
	gormdb "cinema/infra/gormdb"
	"fmt"
	"os"
	"strings"
)

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func main() {
	descricaoGenero := StringPrompt("Qual é seu nome")

	db, err := gormdb.SetupGormDB()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Gender{})

	newGender := entity.Gender{Descricao: descricaoGenero}

	db.Create(&newGender)

	var matchingGender entity.Gender
	// var genders []entity.Gender
	db.First(&matchingGender, newGender.ID)

	fmt.Println(matchingGender.Descricao)

}
