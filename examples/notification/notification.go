package main

import (
	"context"

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
	app.Command("/notificationExampleOne", cmdExampleOne)
	app.Command("/notificationExampleTwo", cmdExampleTwo)

	// start service
	app.Start(apps.AppServerConfig{
		Port:    "8080",
		Timeout: 30, // default 30 second
	})
}

var cmdExampleOne = func(c context.Context) (interface{}, error) {
	// init message component
	newNotif := models.NewNotification(models.NotificationObject{
		Type:       models.MNOTSuccess,
		Title:      "test title",
		Message:    "test body",
		CloseBlock: true,
	})

	// Send is the last step for creating component
	// there is compose, validate component and the result will be send to client
	return newNotif.Send()
}

var cmdExampleTwo = func(c context.Context) (interface{}, error) {
	// init message component
	newNotif := models.NewNotification(models.NotificationObject{
		Type:    models.MNOTFailed,
		Title:   "test title",
		Message: "test body",
	})

	// Send is the last step for creating component
	// there is compose, validate component and the result will be send to client
	return newNotif.Send()
}
