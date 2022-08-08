package apps

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerCommand func(*gin.Context) (interface{}, error)

func (ths *app) Command(commandName string, commanHandler HandlerCommand) {
	command := ths.EngineApps.Group("/command")

	handler := func(c *gin.Context) {
		result, err := commanHandler(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, "error")
		}
		c.JSON(http.StatusOK, result)
	}
	command.POST(commandName, handler)
}
