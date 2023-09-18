package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StringPrompt(label string) string {
	var s string

	if len(label) > 0 {
		fmt.Println(label)
	}

	r := bufio.NewReader(os.Stdin)
	for {
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func CrudMenuPrompt(label string) string {
	fmt.Printf("\n  Menu de %s. Escolha uma opção:\n", label)
	fmt.Println("    [1] Listar todos")
	fmt.Println("    [2] Criar")
	fmt.Println("    [3] Editar")
	fmt.Println("    [4] Excluir")
	fmt.Printf("    [5] Voltar\n\n")

	input := StringPrompt("")

	return input
}