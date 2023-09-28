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

func IntPrompt(label string) uint {
	var i uint

	if len(label) > 0 {
		fmt.Println(label)
	}

	_, err := fmt.Scanln(&i)
	if err != nil {
		fmt.Println("Erro ao ler valor")
	}

	return i
}

func CrudMenuPrompt(label string) string {
	fmt.Printf("\nMenu de %s. Escolha uma opção:\n", label)
	fmt.Println("[1] Listar todos")
	fmt.Println("[2] Criar")
	fmt.Println("[3] Editar")
	fmt.Println("[4] Excluir")
	fmt.Printf("[5] Voltar\n\n")

	input := StringPrompt("")
	PrintDivider()
	return input
}

const (
	ConfirmOption = "s"
	CancelOption  = "n"
)

func ConfirmDeletePrompt() bool {
	for {
		input := StringPrompt("Confirmar exclusão? (s)im/(n)ão")

		switch input {
		case ConfirmOption:
			return true
		case CancelOption:
			return false
		default:
			PrintInvalidOption()
		}
	}
}

func PrintDivider() {
	fmt.Println("----------------------------------------------------------")
}

func PrintDiscardedChanges() {
	fmt.Println("Alterações descartadas.")
}

func PrintInvalidOption() {
	fmt.Println("Opção inválida.")
}
