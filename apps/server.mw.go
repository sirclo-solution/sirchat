package apps

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirclo-solution/sirchat/logger"
)

type key int

const (
	reqBody key = iota
	sircloAuth
)

// method for verifying request using HMAC and SHA256
func verifyingRequest(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			ResponseError(c, NewAppsError(http.StatusBadRequest, err, "bad request"))
			return
		}
		if sirchatSignature := c.GetHeader("X-Sirchat-Signature"); sirchatSignature != "" {
			if ok, err := VerifySignatureSirchat(requestBody, secretKey, sirchatSignature); !ok {
				if err != nil {
					logger.Get().ErrorWithoutSTT("Verify Signature Sirchat", "Error", err)
				}
				if err == nil {
					err = fmt.Errorf("invalid signature")
				}
				ResponseError(c, NewAppsError(http.StatusUnauthorized, err, "invalid signature"))
				return
			}
		} else {
			logger.Get().ErrorWithoutSTT("Header Sirchat Signature", "Error", "signature is required")
			err = errors.New("signature is required")
			ResponseError(c, NewAppsError(http.StatusUnauthorized, err, "signature is required"))
			return
		}

		// set request body to Sirchat-Request-Body
		ctx := context.WithValue(c.Request.Context(), reqBody, requestBody)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// method for forwarding authorization sirclo (only use internal sirclo)
func forwardingSircloAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// only use internal Sirclo
		if authSirclo := c.GetHeader("X-Sirclo-Authorization"); authSirclo != "" {
			ctx := context.WithValue(c.Request.Context(), sircloAuth, authSirclo)
			c.Request = c.Request.WithContext(ctx)
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
