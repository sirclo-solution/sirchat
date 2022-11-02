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
	app.Command("/drawerExampleCard", cmdExampleCard)

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

	// NotificationBlockObject holds the detail of notification block
	newNotificationBlock := models.NotificationBlockObject{
		Title:   "This is a bunch of useful information",
		Message: "Information notification, write some useful notification here.",
		Type:    models.NotificationBlockTypeInfo,
	}
	// NewNotificationBlock used to create new notification block
	notificationBlock := models.NewNotificationBlock(newNotificationBlock)

	// AddBlocks on component for creating Block for wrapping all the blocks
	newDrawer.AddBlocks(notificationBlock)

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

	newSuccessNotificationBlock := models.NotificationBlockObject{
		Title:   "This is an example of success Notification",
		Message: "Success notification, write information thath contains the success of an action.",
		Type:    models.NotificationBlockTypeSuccess,
	}
	// NewNotificationBlock used to create new notification block
	notificationBlock := models.NewNotificationBlock(newSuccessNotificationBlock)

	// AddBlocks on component for creating Block for wrapping all the blocks
	newDrawer.AddBlocks(notificationBlock)

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

	iconBlock := models.NewIconBlock(models.IconBlockObject{
		Src: "https://example.com/dummy1.jpg",
		Alt: "a dummy icon",
	})

	textBlock2 := models.NewTextBlock(&models.TextBlockObject{
		Body:     "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		ViewMore: true,
		Min:      100,
	})

	// example for add new block on container block
	containerBlock.Container.AddBlocks(iconBlock, textBlock2)

	newCollapsibleBlock := models.NewCollapsibleBlock(models.CollapsibleBlockObject{
		Title:     "title collapsible",
		Collapsed: true,
	})

	newCollapsibleBlock.AddContentCollapsible(textBlock2, newbutton)

	// AddBlocks on component for creating Block for wrapping all the blocks
	newDrawer.AddBlocks(containerBlock, newCollapsibleBlock)

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

var cmdExampleCard = func(c context.Context) (interface{}, error) {
	// init drawer component
	newDrawer := models.NewDrawer()

	// NewTitle for adding block title
	newDrawer.Title = models.NewTitle(models.Title{
		Text: "Riwayat Pesanan",
	})

	textBlockHeader := models.NewTextBlock(&models.TextBlockObject{
		Body: "12 Agustus 2022",
	})

	// to be update using pill block
	pillBlock := models.NewTextBlock(&models.TextBlockObject{
		Body: "PESANAN SELESAI",
	})

	// NewContainerBlock use for creating new container block
	// in container block can embed/append another block
	containerBlock := models.NewContainerBlock(models.ContainerBlockObject{
		Direction: models.CDRow,
	})

	// example for add new block on container block
	containerBlock.Container.AddBlocks(textBlockHeader, pillBlock)

	textBlockNoPesanan := models.NewTextBlock(&models.TextBlockObject{
		Body: "No. Pesanan: 312352347128090000",
	})

	textBlockInvoice := models.NewTextBlock(&models.TextBlockObject{
		Body: "INV/20220824/MPL/22214124",
	})

	submitButton := models.NewButtonBlock(models.ButtonBlockObject{
		Type:  models.MBTTSubmit,
		Label: "Lihat Detail Pesanan",
	})

	// NewAction is action from the button
	// add buttons when creating the Action object
	newDrawer.Action = models.NewAction("lihatPesanan")

	newCardBlock := models.NewCardBlock(models.CardBlockObject{})
	newCardBlock.AddCardHeader(containerBlock)
	newCardBlock.Card.AddBlocks(textBlockNoPesanan, textBlockInvoice, submitButton)

	// AddBlocks on component for creating Block for wrapping all the blocks
	newDrawer.AddBlocks(newCardBlock)

	// Send is the last step for creating component
	// there is compose, validate component and the result will be send to client
	return newDrawer.Send()
}
