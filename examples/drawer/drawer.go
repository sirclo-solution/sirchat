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
	newDrawer.Title = models.NewTitle("Drawer Example One", "https://source.unsplash.com/random/50x50")

	// NewCancelButton is button cancel
	cancelButton := models.NewCancelButton("tutup")

	// NewSubmitButton is button submit to the next process/command
	// the action get from first param on NewAction
	submitButton := models.NewSubmitButton("lanjutkan")

	// NewAction is action from the button
	// add buttons when creating the Action object
	newDrawer.Action = models.NewAction("updateCartItems")

	// AddButtons is method for field buttons
	newDrawer.Action.AddButtons(cancelButton, submitButton)

	// add new block on table block
	block1 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 1"})
	block2 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 2"})
	block3 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 3"})
	block4 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 4"})

	// add header for table block
	tableHeader1 := models.HeaderObject{
		Type: "header",
		Text: &models.TextHeaderObject{Align: "horizontal", Body: "Kolom Satu"},
	}
	tableHeader2 := models.HeaderObject{
		Type: "header",
		Text: &models.TextHeaderObject{Align: "horizontal", Body: "Kolom Dua"},
	}
	tableHeaders := []models.HeaderObject{tableHeader1, tableHeader2}

	// add row on table block
	row1 := [][]models.IBlock{
		{
			block1,
		},
		{
			block2,
		},
	}
	row2 := [][]models.IBlock{
		{
			block3,
		},
		{
			block4,
		},
	}

	tableRows := [][][]models.IBlock{row1, row2}

	// NewTableBlock use createng new table block
	table := models.NewTableBlock(
		tableHeaders,
		tableRows,
	)

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
	newDrawer.Title = models.NewTitle("Drawer Example One", "https://source.unsplash.com/random/50x50")

	// NewCancelButton is button cancel
	cancelButton := models.NewCancelButton("tutup")

	// NewSubmitButton is button submit to the next process/command
	// the action get from first param on NewAction
	submitButton := models.NewSubmitButton("lanjutkan")

	// NewAction is action from the button
	// add buttons when creating the Action object
	newDrawer.Action = models.NewAction("updateCartItems")

	// AddButtons is method for field buttons
	newDrawer.Action.AddButtons(cancelButton, submitButton)

	// add new block on table block
	block1 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 1"})
	block2 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 2"})
	block3 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 3"})
	block4 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 4"})

	// add header for table block
	tableHeader1 := models.HeaderObject{
		Type: "header",
		Text: &models.TextHeaderObject{Align: "horizontal", Body: "Kolom Satu"},
	}
	tableHeader2 := models.HeaderObject{
		Type: "header",
		Text: &models.TextHeaderObject{Align: "horizontal", Body: "Kolom Dua"},
	}
	tableHeaders := []models.HeaderObject{tableHeader1, tableHeader2}

	// add row on table block
	row1 := [][]models.IBlock{
		{
			block1,
		},
		{
			block2,
		},
	}
	row2 := [][]models.IBlock{
		{
			block3,
		},
		{
			block4,
		},
	}

	tableRows := [][][]models.IBlock{row1, row2}

	// NewTableBlock use createng new table block
	table := models.NewTableBlock(
		tableHeaders,
		tableRows,
	)

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
