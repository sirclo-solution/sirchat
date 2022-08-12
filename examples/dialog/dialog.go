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
	newDialog.Title = models.NewTitle("Dialog Example One", "https://source.unsplash.com/random/50x50")

	query := map[string]interface{}{
		"brandID": "test",
		"cartID":  "123456789",
	}

	// NewActionButton is button have action for next process/command
	actionButton := models.NewActionButton("cari produk", "initSearchProduct", query)

	// NewCancelButton is button cancel
	cancelButton := models.NewCancelButton("tutup")

	// NewSubmitButton is button submit to the next process/command
	// the action get from first param on NewAction
	submitButton := models.NewSubmitButton("lanjutkan")

	// NewAction is action from the button
	// add buttons when creating the Action object
	newDialog.Action = models.NewAction("updateCartItems")

	// AddButtons is method for field buttons
	newDialog.Action.AddButtons(actionButton, cancelButton, submitButton)

	// NewTextBlock use for creating new text block
	textBlock := models.NewTextBlock(&models.TextBlockObject{
		Body: "Cari Produk",
	})

	textBlock2 := models.NewTextBlock(&models.TextBlockObject{
		Body: "a dummy text",
	})

	// NewTextBlock use for creating new image block
	imageBlock := models.NewImageBlock(
		// change to invalid url like "https://example.com/dummy.m4a" to induce error
		"https://example.com/dummy.jpg",
		"a dummy image",
	)

	// NewContainerBlock use for creating new container block
	// in container block can embed/append another block
	containerBlock := models.NewContainerBlock("row")

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
	newDialog.Title = models.NewTitle("Dialog Example Two", "https://source.unsplash.com/random/50x50")

	// NewCancelButton is button cancel
	cancelButton := models.NewCancelButton("tutup")

	// NewSubmitButton is button submit to the next process/command
	// the action get from first param on NewAction
	submitButton := models.NewSubmitButton("lanjutkan")

	// NewAction is action from the button
	// add buttons when creating the Action object
	newDialog.Action = models.NewAction("addToCart")

	// AddButtons is method for field buttons
	newDialog.Action.AddButtons(cancelButton, submitButton)

	// NewTextBlock use for creating new text block
	textBlock := models.NewTextBlock(&models.TextBlockObject{
		Body: "Cari Produk",
	})

	// NewTextBlock use for creating new input block
	inputBlock := models.NewInputBlock(models.InputBlockObjectTypeText)
	inputBlock.Input.Value = "jacket"
	inputBlock.Input.Name = "query"
	inputBlock.Input.Placeholder = "Masukkan nama produk atau SKU"

	// NewTextBlock use for creating new image block
	imageBlock := models.NewImageBlock(
		// change to invalid url like "https://example.com/dummy.m4a" to induce error
		"https://example.com/dummy.jpg",
		"a dummy image",
	)

	// NewContainerBlock use for creating new container block
	// in container block can embed/append another block
	containerBlock := models.NewContainerBlock("row")

	// example for add new block on container block
	containerBlock.Container.AddBlocks(imageBlock, inputBlock)

	// AddBlocks on component for creating Block for wrapping all the blocks
	newDialog.AddBlocks(textBlock, containerBlock)

	// Send is the last step for creating component
	// there is compose, validate component and the result will be send to client
	return newDialog.Send()
}
