package buyercli

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/utils"
	"fmt"
)

const (
	ListAllOption = "1"
	CreateOption  = "2"
	EditOption    = "3"
	DeleteOption  = "4"
	BackOption    = "5"

	EditNameOption     = "1"
	EditDocumentOption = "2"
	EditSaveOption     = "s"
	EditCancelOption   = "c"
	SaveChanges        = true
	DiscardChanges     = false
)

func BuyerEditPrompt(buyer *entities.Buyer) bool {
	PrintBuyerLabel()
	PrintBuyer(*buyer)

	for {
		fmt.Println("Edição de Pagante, escolha uma opção:")
		fmt.Println("[1] Nome")
		fmt.Println("[2] Documento")
		fmt.Println("[s] Salvar")
		fmt.Println("[c] Cancelar")

		input := utils.StringPrompt("")

		switch input {
		case EditNameOption:
			buyer.Name = utils.StringPrompt("Nome do Pagante:")
		case EditDocumentOption:
			buyer.Document = utils.StringPrompt("Documento do Pagante:")
		case EditSaveOption:
			return true
		case EditCancelOption:
			return false
		default:
			utils.PrintInvalidOption()

		}
	}
}

func PrintBuyer(buyer entities.Buyer) {
	fmt.Printf("[%v]: %s | %s\n", buyer.BuyerID, buyer.Name, buyer.Document)
}

func PrintBuyerLabel() {
	fmt.Println("[id]: Nome | Documento")
}
