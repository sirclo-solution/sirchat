package apps

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// key of header signature HMAC for verifying request
	HeaderSirchatSignature = "X-Sirchat-Signature"

	// key of header Authorization Sirclo (its only use internal SIRCLO)
	HeaderSircloAuthorization = "X-Sirclo-Authorization"

	// key of Authorization Sirclo (its only use internal SIRCLO)
	SircloAuthorization = "Sirclo-Authorization"

	// key of get Request Body
	SirchatRequestBody = "Sirchat-Request-Body"
)

// method for verifying request using HMAC and SHA256
func verifyingRequest(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			ResponseError(c, NewAppsError(http.StatusBadRequest, err, "bad request"))
			return
		}
		if sirchatSignature := c.GetHeader(HeaderSirchatSignature); sirchatSignature != "" {
			if ok, err := VerifySignatureSirchat(requestBody, secretKey, sirchatSignature); !ok {
				if err != nil {
					log.Println("Error VerifySignatureSirchat(): ", err)
				}
				if err == nil {
					err = fmt.Errorf("invalid signature")
				}
				ResponseError(c, NewAppsError(http.StatusUnauthorized, err, "invalid signature"))
				return
			}
		} else {
			log.Println("[Sirchat] - signature is required")
			err = errors.New("signature is required")
			ResponseError(c, NewAppsError(http.StatusUnauthorized, err, "signature is required"))
			return
		}

		// set request body to Sirchat-Request-Body
		c.Set(SirchatRequestBody, requestBody)

		c.Next()
	}
}

// method for forwarding authorization sirclo (only use internal sirclo)
func forwardingSircloAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// only use internal Sirclo
		if authSirclo := c.GetHeader(HeaderSircloAuthorization); authSirclo != "" {
			c.Set(SircloAuthorization, authSirclo)
		}
		c.Next()
	}
}

// method for generate signature sirchat
func GenerateSignatureSirchat(body []byte, secretKey string) string {
	hmac := hmac.New(sha256.New, []byte(secretKey))
	hmac.Write(body)

	return hex.EncodeToString(hmac.Sum(nil))
}

// method for verify signature sirchat
func VerifySignatureSirchat(body []byte, secretKey, signature string) (bool, error) {
	sign, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write(body)

	return hmac.Equal(sign, mac.Sum(nil)), nil
}
