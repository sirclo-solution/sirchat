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
	app.Command("/messageExampleOne", cmdExampleOne)

	// start service
	app.Start(apps.AppServerConfig{
		Port:    "8080",
		Timeout: 30, // default 30 second
	})
}

var cmdExampleOne = func(c *gin.Context) (interface{}, error) {
	// init message component
	newMessage := models.NewMessage(models.MessageObject{
		TenantID: "chat",
		BrandID:  "chat",
		RoomID:   "room",
		Channel:  "channel",
	})

	// to be updated using method
	newMessage.Message.Texts = []models.MessageTextObject{
		{
			Body: "message 1",
		},
		{
			Body: "message 2",
		},
	}
	//// to be updated using method
	newMessage.Message.Images = []models.MessageImageObject{
		{
			Alt: "Image 1",
			Src: "https://example.com/dummy.jpg",
		},
		{
			Alt: "Image 2",
			Src: "https://example.com/dummy.jpg",
		},
	}

	// Send is the last step for creating component
	// there is compose, validate component and the result will be send to client
	return newMessage.Send()
}
