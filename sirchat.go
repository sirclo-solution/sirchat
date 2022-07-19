package main

import (
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

	textBlock2 := models.NewTextBlock(&models.TextBlockObject{
		Type: "label",
		Body: "a dummy text",
	})

	imageBlock := models.NewImageBlock(&models.ImageBlockObject{
		Src: "https://example.com/dummy.jpg", // change to invalid url like "https://example.com/dummy.m4a" to induce error
		Alt: "a dummy image",
	})

	containerBlock := models.NewContainerBlock(&models.ContainerBlockObject{
		Direction: "row",
	})

	containerBlock2 := models.NewContainerBlock(&models.ContainerBlockObject{
		Direction: "row",
	})

	containerBlock3 := models.NewContainerBlock(&models.ContainerBlockObject{
		Direction: "row", // change to something like "fake_row" to induce error
	})

	containerBlock3.Container.AddBlock(textBlock2)

	containerBlock2.Container.AddBlock(containerBlock3)

	containerBlock.Container.AddBlock(imageBlock)

	newDialog.Blocks = append(newDialog.Blocks, textBlock, containerBlock, containerBlock2)
	result, errs := newDialog.Compose()
	if errs != nil {
		fmt.Printf("%+q\n", errs)
		return
	}
	fmt.Printf("Result : %v\n", string(result))
}
