package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirclo-solution/sirchat/apps"
	"github.com/sirclo-solution/sirchat/models"
)

const (
	SECRET_KEY = "sirchat-sirclochat"
)

func main() {
	// creating new apps
	app := apps.NewApps(apps.AppConfig{
		AppSecret: SECRET_KEY,
	})

	// creating new action/command/api
	app.Command("/dialogExampleOne", cmdExampleOne)
	app.Command("/dialogExampleTwo", cmdExampleTwo)

	// start service
	app.Start(apps.AppServerConfig{
		Port:    "8080",
		Timeout: 30, // default 30 second
	})
}

var cmdExampleOne = func(c *gin.Context) (interface{}, error) {
	// init dialog component
	newDialog := models.NewDialog()

	// NewTitle for adding block title
	newDialog.Title = models.NewTitle(models.Title{
		Text: "Dialog Example One",
	})

	query := map[string]interface{}{
		"brandID": "test",
		"cartID":  "123456789",
	}

	// NewActionButton is button have action for next process/command
	actionButton := models.NewButtonBlock(models.ButtonBlockObject{
		Type:  models.MBTTAction,
		Label: "Lihat Keranjang",
		Action: &models.ButtonActionObject{
			ID: "viewCart",
		},
		Query: query,
	})

	// NewCancelButton is button cancel
	cancelButton := models.NewButtonBlock(models.ButtonBlockObject{
		Type:  models.MBTTCancel,
		Label: "tutup",
	})

	// NewSubmitButton is button submit to the next process/command
	// the action get from first param on NewAction
	submitButton := models.NewButtonBlock(models.ButtonBlockObject{
		Type:  models.MBTTSubmit,
		Label: "lanjutkan",
	})

	// NewAction is action from the button
	// add buttons when creating the Action object
	newDialog.Action = models.NewAction("updateCartItems")

	// AddButtons is method for field buttons
	newDialog.Action.AddButtons(actionButton.Button, cancelButton.Button, submitButton.Button)

	// NewTextBlock use for creating new text block
	textBlock := models.NewTextBlock(&models.TextBlockObject{
		Body: "Cari Produk",
	})

	textBlock2 := models.NewTextBlock(&models.TextBlockObject{
		Body: "a dummy text",
	})

	// NewTextBlock use for creating new image block
	imageBlock := models.NewImageBlock(models.ImageBlockObject{
		// change to invalid url like "https://example.com/dummy.m4a" to induce error
		Src: "https://example.com/dummy.jpg",
		Alt: "a dummy image",
	})

	// NewContainerBlock use for creating new container block
	// in container block can embed/append another block
	containerBlock := models.NewContainerBlock(models.ContainerBlockObject{
		Direction: models.CDRow,
	})

	// example for add new block on container block
	containerBlock.Container.AddBlocks(imageBlock, textBlock2)

	// AddBlocks on component for creating Block for wrapping all the blocks
	newDialog.AddBlocks(textBlock, containerBlock)

	// Send is the last step for creating component
	// there is compose, validate component and the result will be send to client
	return newDialog.Send()
}

var cmdExampleTwo = func(c *gin.Context) (interface{}, error) {
	// init dialog component
	newDialog := models.NewDialog()

	// NewTitle for adding block title
	newDialog.Title = models.NewTitle(models.Title{
		Text: "Dialog Example Two",
	})

	// NewCancelButton is button cancel
	cancelButton := models.NewButtonBlock(models.ButtonBlockObject{
		Type:  models.MBTTCancel,
		Label: "tutup",
	})

	// NewSubmitButton is button submit to the next process/command
	// the action get from first param on NewAction
	submitButton := models.NewButtonBlock(models.ButtonBlockObject{
		Type:  models.MBTTSubmit,
		Label: "lanjutkan",
	})

	// NewAction is action from the button
	// add buttons when creating the Action object
	newDialog.Action = models.NewAction("addToCart")

	// AddButtons is method for field buttons
	newDialog.Action.AddButtons(cancelButton.Button, submitButton.Button)

	// NewTextBlock use for creating new text block
	textBlock := models.NewTextBlock(&models.TextBlockObject{
		Body: "Cari Produk",
	})

	// NewTextBlock use for creating new input block
	inputBlock := models.NewInputBlock(&models.InputBlockObject{
		Type:        models.InputBlockObjectTypeText,
		Value:       "jacket",
		Name:        "query",
		Placeholder: "Masukkan nama produk atau SKU",
	})

	// NewTextBlock use for creating new image block
	imageBlock := models.NewImageBlock(models.ImageBlockObject{
		// change to invalid url like "https://example.com/dummy.m4a" to induce error
		Src: "https://example.com/dummy.jpg",
		Alt: "a dummy image",
	})

	// NewContainerBlock use for creating new container block
	// in container block can embed/append another block
	containerBlock := models.NewContainerBlock(models.ContainerBlockObject{
		Direction: "row",
	})

	// example for add new block on container block
	containerBlock.Container.AddBlocks(imageBlock, inputBlock)

	textListBlock := models.NewTextListBlock()
	for i := 1; i <= 3; i++ {
		textListBlock.AddTextBlock(textBlock)
	}

	// AddBlocks on component for creating Block for wrapping all the blocks
	newDialog.AddBlocks(textBlock, containerBlock, textListBlock)

	// Send is the last step for creating component
	// there is compose, validate component and the result will be send to client
	return newDialog.Send()
}
