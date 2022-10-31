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

	// NewLinkButton is button to open requested url
	url := "https://sirclo.com"
	linkButton := models.NewButtonBlock(models.ButtonBlockObject{
		Type:  models.MBTTAction,
		Label: "Home",
		Action: &models.ButtonActionObject{
			ID:   "",
			Link: &url,
		},
	})

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
	newDialog.Action.AddButtons(linkButton.Button, actionButton.Button, cancelButton.Button, submitButton.Button)

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

	newNotif := models.NewNotification(models.NotificationObject{
		Type:    models.MNOTSuccess,
		Title:   "Ini notifikasi sukses",
		Message: "Success broo",
	})

	newDialog.Notification = &newNotif.Notification

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

	minInputText := 1

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
		MinInput:    &minInputText,
		MaxInput:    &minInputText,
	})

	inputBlockRadio := models.NewInputBlock(&models.InputBlockObject{
		Type:        models.InputBlockObjectTypeDistrictSelect,
		Value:       "jacket",
		Name:        "query",
		Placeholder: "Masukkan nama produk atau SKU",
		Required:    new(bool),
		MinInput:    &minInputText,
		MaxInput:    &minInputText,
	})

	// NewTextBlock use for creating new image block
	imageBlock := models.NewImageBlock(models.ImageBlockObject{
		// change to invalid url like "https://example.com/dummy.m4a" to induce error
		Src: "https://example.com/dummy.jpg",
		Alt: "a dummy image",
	})

	// NewTextBlock use for creating new image block
	imageBlock2 := models.NewImageBlock(models.ImageBlockObject{
		// change to invalid url like "https://example.com/dummy.m4a" to induce error
		Src: "https://storage.googleapis.com/sirclo-1152-storefront/products/fd6bcc6d-506e-4e14-a530-b1ea55c33fd6-bajaksawah%.jpg",
		Alt: "a dummy image",
	})

	// NewTextBlock use for creating new image block
	imageBlock3 := models.NewImageBlock(models.ImageBlockObject{
		// change to invalid url like "https://example.com/dummy.m4a" to induce error
		Src: "https://storage.googleapis.com/sirclo-1152-storefront/products/53628732-167b-4938-955e-7da5c1c010ca-17 agustus.jpeg",
		Alt: "a dummy image",
	})

	// NewTextBlock use for creating new image block
	imageBlock4 := models.NewImageBlock(models.ImageBlockObject{
		// change to invalid url like "https://example.com/dummy.m4a" to induce error
		Src: "https://storage.googleapis.com/sirclo-1152-storefront/products/53628732-167b-4938-955e-7da5c1c&%34010ca-17 agustus.png",
		Alt: "a dummy image",
	})

	// NewTextBlock use for creating new image block
	imageBlockWithoutExt := models.NewImageBlock(models.ImageBlockObject{
		// change to invalid url like "https://example.com/dummy.m4a" to induce error
		Src: "https://sirclocdn.com/chat/products/_210323141909_2",
		Alt: "a dummy image",
	})

	inputCounter := models.NewInputBlock(&models.InputBlockObject{
		Type:  models.InputBlockObjectTypeCounter,
		Value: "1",
		Name:  "counter1",
	})

	numberOnly := true
	maxInputNumber := 20
	inputNumber := models.NewInputBlock(&models.InputBlockObject{
		Type:       models.InputBlockObjectTypeNumber,
		Value:      "1",
		Name:       "number",
		MaxInput:   &maxInputNumber,
		NumberOnly: &numberOnly,
	})

	minInput := 1
	maxInput := 10
	inputCounter2 := models.NewInputBlock(&models.InputBlockObject{
		Type:     models.InputBlockObjectTypeCounter,
		Value:    "1",
		Name:     "counter2",
		MinInput: &minInput,
		MaxInput: &maxInput,
	})

	inputBlockNumberOnly := models.NewInputBlock(&models.InputBlockObject{
		Type:        models.InputBlockObjectTypeText,
		Value:       "jacket",
		Name:        "query",
		Placeholder: "Masukkan nama produk atau SKU",
		Required:    new(bool),
		MinInput:    &minInputText,
		MaxInput:    &minInputText,
		NumberOnly:  &numberOnly,
	})

	textTitleBlock := models.NewTextBlock(&models.TextBlockObject{
		Type:  models.TextBlockObjectTypeTitle,
		Body:  "Ini Title",
		Color: models.TextBlockObjectColorPrimary,
	})

	// NewContainerBlock use for creating new container block
	// in container block can embed/append another block
	containerBlock := models.NewContainerBlock(models.ContainerBlockObject{
		Direction: "row",
	})

	// example for add new block on container block
	containerBlock.Container.AddBlocks(imageBlock, imageBlock2, imageBlock3, imageBlock4, imageBlockWithoutExt, inputBlock, inputCounter, inputNumber, inputCounter2, inputBlockRadio, inputBlockNumberOnly, textTitleBlock)

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
