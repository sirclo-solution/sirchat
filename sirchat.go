package main

import (
	"encoding/json"
	"fmt"

	"github.com/sirclo-solution/sirchat/models"
)

func main() {
	fmt.Println("service is running ...")
	app := models.NewApp()
	newDialog := app.NewDialog()
	newDialog.Title = models.NewTitle("ini text", "ini icon")
	actionButton := models.NewActionButton("cari produk", "initSearchProduct")
	cancelButton := models.NewCancelButton("tutup")
	submitButton := models.NewSubmitButton("lanjutkan")
	newDialog.Action = models.NewAction("updateCartItems", models.NewButtons(actionButton, cancelButton, submitButton))
	textBlock := models.NewTextBlock(&models.TextBlockObject{
		Type: "text",
		Body: "Cari Produk",
	})
	newDialog.Blocks = textBlock
	result, err := json.Marshal(newDialog)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result : %v\n", string(result))
}
