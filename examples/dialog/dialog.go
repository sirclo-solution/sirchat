package main

import (
	"context"

	"github.com/sirclo-solution/sirchat/apps"
	"github.com/sirclo-solution/sirchat/logger"
	"github.com/sirclo-solution/sirchat/models"
)

const (
	SECRET_KEY = "dummy-key"
)

type RequestExampleOne struct {
	Chat      ChatDetail `json:"chat"`
	PayloadSP *Payload   `json:"payload"`
}

type ChatDetail struct {
	TenantId string `json:"tenantId"`
	BrandId  string `json:"brandId"`
	RoomId   string `json:"roomId"`
	Channel  string `json:"channel"`
}

type Payload struct {
	Query string `json:"query"`
}

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

var cmdExampleOne = func(c context.Context) (interface{}, error) {
	// get auth sirclo (only use internal sirclo)
	authSirclo, err := apps.GetAuthSirclo(c)
	if err != nil {
		logger.Get().Error("[dialogExampleOne] - Error GetAuthSirclo", "Error", err)
		return nil, err
	}

	logger.Get().Info("[dialogExampleOne] - Auth Sirclo", "Info", authSirclo)

	var req RequestExampleOne

	// bind request  body
	if err := apps.BindRequestBody(c, &req); err != nil {
		logger.Get().Error("[dialogExampleOne] - Error Bind Request Body", "Error", err)
		return nil, err
	}

	logger.Get().Info("[dialogExampleOne] - Request body", "Info", req)

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

var cmdExampleTwo = func(c context.Context) (interface{}, error) {
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

	// NewTextBlock use for creating new input block.
	// Field `Required`` is default to true. You can set it to
	// false by passing a pointer to false value. You can either
	// create a boolean variable and then pass the pointer or directly
	// pass pointer of boolean with `new(bool)`
	inputBlock := models.NewInputBlock(&models.InputBlockObject{
		Type:        models.InputBlockObjectTypeText,
		Value:       "jacket",
		Name:        "query",
		Placeholder: "Masukkan nama produk atau SKU",
		Required:    new(bool),
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
