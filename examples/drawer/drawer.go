package main

import (
	"context"
	"strconv"

	"github.com/sirclo-solution/sirchat/apps"
	"github.com/sirclo-solution/sirchat/models"
)

const (
	SECRET_KEY = "dummy-key"
)

func main() {
	// creating new apps
	app := apps.NewApps(apps.AppConfig{
		AppSecret: SECRET_KEY,
	})

	// creating new action/command/api
	app.Command("/drawerExampleOne", cmdExampleOne)
	app.Command("/drawerExampleTwo", cmdExampleTwo)
	app.Command("/drawerExampleThree", cmdExampleThree)

	// start service
	app.Start(apps.AppServerConfig{
		Port:    "8080",
		Timeout: 30, // default 30 second
	})
}

var cmdExampleOne = func(c context.Context) (interface{}, error) {
	// init drawer component
	newDrawer := models.NewDrawer()

	// NewTitle for adding block title
	newDrawer.Title = models.NewTitle(models.Title{
		Text: "Drawer Example One",
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
	newDrawer.Action = models.NewAction("updateCartItems")

	// AddButtons is method for field buttons
	newDrawer.Action.AddButtons(cancelButton.Button, submitButton.Button)

	// add new block on table block
	block1 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 1"})
	block2 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 2"})
	block3 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 3"})
	block4 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 4"})

	// add new input block type radio
	newInputRadio := models.NewInputBlock(&models.InputBlockObject{
		Type: models.InputBlockObjectTypeRadio,
		Name: "shippingMethod",
	})

	// input radio should be have minimal 1 options.
	inputRadioOption := models.InputBlockOptionsObject{
		Value: "JNE-REG",
		Label: "JNE Reguler",
	}

	// use AddDescriptions to add array of text block as the descriptions for each option.
	// when the input is "select" type, the descriptions will not be rendered in the UI.
	// parameter Descriptions is optional
	inputRadioOption.AddDescriptions(
		*models.NewTextBlock(&models.TextBlockObject{Body: "1-2 Hari"}),
		*models.NewTextBlock(&models.TextBlockObject{Body: "P. Jawa dan luar P. Jawa"}),
	)

	// AddInputBlockOptionsObject use to be add options on input radio.
	newInputRadio.AddInputBlockOptionsObject(inputRadioOption)

	// NewTableBlock use createng new table block
	table := models.NewTableBlock()

	// add header on table block
	for i := 1; i <= 2; i++ {
		table.AddHeader(models.TextHeaderObject{
			Body: "header" + strconv.Itoa(i),
		})
	}

	// rows is amount of data
	var rows [][][]models.IBlock
	for i := 1; i <= 2; i++ {
		// columns is the data held by each row.
		// each column can be filled more than 1 block.
		// the number of headers and columns must be the same
		var columns [][]models.IBlock
		for j := 1; j <= 2; j++ {
			column := table.AddColumn(block1, block2, newInputRadio)
			if j == 2 {
				column = table.AddColumn(block3, block4)
			}
			columns = append(columns, column)
		}
		row := table.AddRow(columns...)
		rows = append(rows, row)
	}

	// add body on table block
	table.AddBody(rows...)

	// AddBlocks on component for creating Block for wrapping all the blocks
	newDrawer.AddBlocks(table)

	// Send is the last step for creating component
	// there is compose, validate component and the result will be send to client
	return newDrawer.Send()
}

var cmdExampleTwo = func(c context.Context) (interface{}, error) {
	// init drawer component
	newDrawer := models.NewDrawer()

	// NewTitle for adding block title
	newDrawer.Title = models.NewTitle(models.Title{
		Text: "Drawer Example Two",
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

	// NewIconButton is button icon
	// this button has an action when clicked it will trigger to the next action
	iconButton := models.NewButtonBlock(models.ButtonBlockObject{
		Type: models.MBTTAction,
		Icon: models.ButtonObjectIconCart,
		Action: &models.ButtonActionObject{
			ID: "addToCart",
		},
	})

	// NewAction is action from the button
	// add buttons when creating the Action object
	newDrawer.Action = models.NewAction("updateCartItems")

	// AddButtons is method for field buttons
	newDrawer.Action.AddButtons(cancelButton.Button, submitButton.Button, iconButton.Button)

	// add new block on table block
	block1 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 1"})
	block2 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 2"})
	inputCounter := models.NewInputBlock(&models.InputBlockObject{
		Type:    models.InputBlockObjectTypeCounter,
		Value:   "1",
		Name:    "tas",
		GroupID: "addItem",
		Action: &models.InputActionObject{
			ID: "updateCartItem",
		},
	})
	block3 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 3"})
	block4 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 4"})

	// NewTableBlock use createng new table block
	table := models.NewTableBlock()

	// add header on table block
	for i := 1; i <= 2; i++ {
		table.AddHeader(models.TextHeaderObject{
			Body: "header" + strconv.Itoa(i),
		})
	}

	// rows is amount of data
	var rows [][][]models.IBlock
	for i := 1; i <= 2; i++ {
		// columns is the data held by each row
		// each column can be filled more than 1 block
		// the number of headers and columns must be the same
		var columns [][]models.IBlock
		for j := 1; j <= 2; j++ {
			column := table.AddColumn(block1, block2, inputCounter)
			if j == 2 {
				column = table.AddColumn(block3, block4)
			}
			columns = append(columns, column)
		}
		row := table.AddRow(columns...)
		rows = append(rows, row)
	}

	// add body on table block
	table.AddBody(rows...)

	// NewCarouselBlock use creating new carousel block
	blockCarousel := models.NewCarouselBlock("Title Carousel")
	blockCarousel.AddDescriptionsCarousel("description 1")
	blockCarousel.AddDescriptionsCarousel("description 2")
	blockCarousel.AddImageCarousel(models.ImageBlockObject{
		Alt: "image 1",
		Src: "https://example.com/dummy1.jpg",
	})
	blockCarousel.AddImageCarousel(models.ImageBlockObject{
		Alt: "image 2",
		Src: "https://example.com/dummy2.jpg",
	})

	// AddBlocks on component for creating Block for wrapping all the blocks
	newDrawer.AddBlocks(table, blockCarousel)

	// Send is the last step for creating component
	// there is compose, validate component and the result will be send to client
	return newDrawer.Send()
}

var cmdExampleThree = func(c context.Context) (interface{}, error) {
	// init drawer component
	newDrawer := models.NewDrawer()

	// NewTitle for adding block title
	newDrawer.Title = models.NewTitle(models.Title{
		Text: "Drawer Example Three",
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
		Label: "submit",
	})

	textBlock := models.NewTextBlock(&models.TextBlockObject{
		Body: "Kamu belum menyimpan catatan produk. Perubahan akan terhapus jika kamu keluar dari halaman ini. Apakah kamu yakin untuk melanjutkan?",
	})

	// NewPromptBlock is used to create prompt block
	promptBlock := models.NewPromptBlock(models.PromptBlockObject{
		Title:          "Perubahan Belum Disimpan",
		CancelButton:   &cancelButton.Button,
		ContinueButton: &submitButton.Button,
	})

	promptBlock.AddBlocks(textBlock)

	newbutton := models.NewButtonBlock(models.ButtonBlockObject{
		Type:   models.MBTTSubmit,
		Label:  "lanjutkan",
		Prompt: promptBlock,
	})

	// NewAction is action from the button
	// add buttons when creating the Action object
	newDrawer.Action = models.NewAction("saveItems")

	// AddButtons is method for field buttons
	newDrawer.Action.AddButtons(newbutton.Button)

	// AddOnClose is method for field on_close
	newDrawer.Action.AddOnClose(&models.ActionOnClose{
		Prompt: promptBlock,
	})

	// NewContainerBlock use for creating new container block
	// in container block can embed/append another block
	containerBlock := models.NewContainerBlock(models.ContainerBlockObject{
		Direction: models.CDRow,
	})

	textBlock2 := models.NewTextBlock(&models.TextBlockObject{
		Body: "dummy text",
	})

	// example for add new block on container block
	containerBlock.Container.AddBlocks(textBlock2)

	// AddBlocks on component for creating Block for wrapping all the blocks
	newDrawer.AddBlocks(containerBlock)

	newNotif := models.NewNotification(models.NotificationObject{
		Type:    models.MNOTSuccess,
		Title:   "Ini notifikasi sukses",
		Message: "Success broo",
	})

	newDrawer.Notification = &newNotif.Notification

	// Send is the last step for creating component
	// there is compose, validate component and the result will be send to client
	return newDrawer.Send()
}
