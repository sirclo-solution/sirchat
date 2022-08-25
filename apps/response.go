package apps

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirclo-solution/sirchat/logger"
)

// AppsError is standard object error Sirchat
type AppsError struct {
	// struct for error component
	AppsErr ErrDetail `json:"error"`
}

type ErrDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

// Return detail error message
func (err AppsError) Error() string {
	if err.AppsErr.Err != nil {
		return err.AppsErr.Err.Error()
	}
	return err.AppsErr.Message
}

// generate new apps error
func NewAppsError(code int, err error, message string) error {
	var errMsg string

	if err != nil {
		errMsg = err.Error()
	}

	if message != "" {
		errMsg = message
	}

	return AppsError{
		AppsErr: ErrDetail{
			Code:    code,
			Message: errMsg,
			Err:     err,
		},
	}
}

// resonse error
func ResponseError(c *gin.Context, err error) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Content-Type", "Application/json")
	var ae AppsError
	if errors.As(err, &ae) {
		logger.Get().ErrorWithoutSTT("Response AppsError", "Error", ae.AppsErr.Err.Error())
		c.AbortWithStatusJSON(ae.AppsErr.Code, ae)
		return
	}

	// handle non AppsError types
	logger.Get().ErrorWithoutSTT("Response Error", "Error", err.Error())
	c.AbortWithStatusJSON(http.StatusInternalServerError, AppsError{
		AppsErr: ErrDetail{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		},
	})
}

// response success
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Content-Type", "Application/json")
	c.JSON(http.StatusOK, data)
}
