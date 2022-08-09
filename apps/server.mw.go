package apps

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
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
)

// method for verifying request using HMAC and SHA256
func verifyingRequest(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, err := c.GetRawData()
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
					err = fmt.Errorf("%v invalid signature", sirchatSignature)
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
	log.Println("verify hmac", hex.EncodeToString(mac.Sum(nil))) // for testing

	return hmac.Equal(sign, mac.Sum(nil)), nil
}
