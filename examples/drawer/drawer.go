package main

import (
	"strconv"

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
	app.Command("/drawerExampleOne", cmdExampleOne)
	app.Command("/drawerExampleTwo", cmdExampleTwo)

	// start service
	app.Start(apps.AppServerConfig{
		Port:    "8080",
		Timeout: 30, // default 30 second
	})
}

var cmdExampleOne = func(c *gin.Context) (interface{}, error) {
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
	// AddInputBlockOptionsObject use to be add options on input radio.
	// parameter Description is optional
	newInputRadio.AddInputBlockOptionsObject(models.InputBlockOptionsObject{
		Value:       "JNE-REG",
		Label:       "JNE Reguler",
		Description: "1-2 hari",
	})

	// NewTableBlock use createng new table block
	table := models.NewTableBlock()

	// add header on table block
	for i := 1; i <= 2; i++ {
		table.AddHeader("header"+strconv.Itoa(i), "")
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

var cmdExampleTwo = func(c *gin.Context) (interface{}, error) {
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
	block3 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 3"})
	block4 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 4"})

	// NewTableBlock use createng new table block
	table := models.NewTableBlock()

	// add header on table block
	for i := 1; i <= 2; i++ {
		table.AddHeader("header"+strconv.Itoa(i), "")
	}

	// rows is amount of data
	var rows [][][]models.IBlock
	for i := 1; i <= 2; i++ {
		// columns is the data held by each row
		// each column can be filled more than 1 block
		// the number of headers and columns must be the same
		var columns [][]models.IBlock
		for j := 1; j <= 3; j++ {
			column := table.AddColumn(block1, block2)
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
	blockCarousel.AddImageCarousel("image 1", "https://example.com/dummy1.jpg")
	blockCarousel.AddImageCarousel("image 2", "https://example.com/dummy2.jpg")

	// AddBlocks on component for creating Block for wrapping all the blocks
	newDrawer.AddBlocks(table, blockCarousel)

	// Send is the last step for creating component
	// there is compose, validate component and the result will be send to client
	return newDrawer.Send()
}
