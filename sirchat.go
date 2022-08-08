package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirclo-solution/sirchat/apps"
	"github.com/sirclo-solution/sirchat/models"
)

const (
	SECRET_KEY = "sirchat-sirclochat"
)

func main() {
	app := apps.NewApps(apps.AppConfig{
		AppSecret: SECRET_KEY,
	})

	app.Command("/firstEx", cmdFirstEx)

	app.Command("/searchProduct", func(c *gin.Context) (interface{}, error) {
		//get auth sirclo
		auth := c.GetString(apps.SircloAuthorization)
		fmt.Println(auth)
		return exInitSearchProduct()
	})

	app.Command("/initTable", cmdInitTable)

	app.Start(apps.ParamRunServer{
		Port:    "8080",
		Timeout: 30, // default 30 second
	})
}

var cmdFirstEx = func(c *gin.Context) (interface{}, error) {
	//get auth sirclo
	auth := c.GetString(apps.SircloAuthorization)
	fmt.Println(auth)
	return firstEx()
}

var cmdInitTable = func(c *gin.Context) (interface{}, error) {
	//get auth sirclo
	auth := c.GetString(apps.SircloAuthorization)
	fmt.Println(auth)
	return exInitTable()
}

func firstEx() (interface{}, error) {
	newDialog := models.NewDialog()
	newDialog.Title = models.NewTitle("ini text", "ini icon")

	query := map[string]interface{}{
		"brandID": "test",
		"cartID":  "123456789",
	}
	actionButton := models.NewActionButton("cari produk", "initSearchProduct", query)

	cancelButton := models.NewCancelButton("tutup")
	submitButton := models.NewSubmitButton("lanjutkan")
	newDialog.Action = models.NewAction("updateCartItems", models.NewButtons(actionButton, cancelButton, submitButton)...)

	textBlock := models.NewTextBlock(&models.TextBlockObject{
		Body: "Cari Produk",
	})

	textBlock2 := models.NewTextBlock(&models.TextBlockObject{
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
	return newDialog.Send()
}

func exInitSearchProduct() (interface{}, error) {
	newDialog := models.NewDialog()
	newDialog.Title = models.NewTitle("Cari Produk", "https://source.unsplash.com/random/50x50")

	query := map[string]interface{}{
		"brandID": "test",
		"cartID":  "123456789",
	}
	actionButton := models.NewActionButton("Lihat Keranjang", "viewCart", query)

	cancelButton := models.NewCancelButton("tutup")
	newDialog.Action = models.NewAction("searchProduct", models.NewButtons(actionButton, cancelButton)...)

	textBlock := models.NewTextBlock(&models.TextBlockObject{
		Body: "Cari Produk",
	})

	inputBlock := models.NewInputBlock(&models.InputBlockObject{
		Type:        "text",
		Value:       "jacket",
		Name:        "query",
		Placeholder: "Masukkan nama produk atau SKU",
	})

	containerBlock := models.NewContainerBlock(&models.ContainerBlockObject{
		Direction: "row",
	})

	containerBlock.Container.AddBlock(inputBlock)

	newDialog.Blocks = models.NewBlocks(textBlock, containerBlock)
	return newDialog.Send()
}

func exInitTable() (interface{}, error) {
	drawer := models.NewDrawer()
	drawer.Title = models.NewTitle("Contoh table", "https://source.unsplash.com/random/50x50")

	block1 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 1"})
	block2 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 2"})
	block3 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 3"})
	block4 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 4"})
	block5 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 5"})
	block6 := models.NewTextBlock(&models.TextBlockObject{Body: "isi 6"})

	tableHeader1 := models.HeaderObject{
		Type: "header",
		Text: &models.TextHeaderObject{Align: "horizontal", Body: "Kolom Satu"},
	}
	tableHeader2 := models.HeaderObject{
		Type: "header",
		Text: &models.TextHeaderObject{Align: "horizontal", Body: "Kolom Dua"},
	}
	tableHeaders := []models.HeaderObject{tableHeader1, tableHeader2}

	row1 := [][]models.IBlock{
		{
			block1, block2,
		},
		{
			block3, block4,
		},
	}
	row2 := [][]models.IBlock{
		{
			block5,
		},
		{
			block6,
		},
	}

	tableRows := [][][]models.IBlock{row1, row2}

	table := models.NewTableBlock(
		tableHeaders,
		tableRows,
	)
	drawer.Blocks = models.NewBlocks(table)
	return drawer.Send()
}
