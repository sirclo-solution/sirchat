package apps

import (
	"github.com/gin-gonic/gin"
)

// this is type for method Handler Command
type HandlerCommand func(*gin.Context) (interface{}, error)

// app command for handling route on group /command
func (ths *app) Command(commandName string, commanHandler HandlerCommand) {
	command := ths.EngineApps.Group("/command")

	handler := func(c *gin.Context) {
		result, err := commanHandler(c)
		if err != nil {
			ResponseError(c, err)
			return
		}
		ResponseSuccess(c, result)
	}

	command.POST(commandName, handler)
}
