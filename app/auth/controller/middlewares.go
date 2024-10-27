package api

import (
	"fmt"
	"net/http"
	"portfolio/services/infrastructure/log"

	service "portfolio/services"

	"portfolio/services/infrastructure/policy"
	"portfolio/services/infrastructure/request"

	"github.com/gin-gonic/gin"

	sentrygin "github.com/getsentry/sentry-go/gin"
)

var host string
var port string

func Init(h, p string) {
	host = h
	port = p
}

func can(action string) gin.HandlerFunc {
	client := &policy.Client{
		Addr: fmt.Sprintf("%s:%s", host, port),
	}
	return func(ctx *gin.Context) {
		userID := request.GetUserID(ctx.Request)

		if userID == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		allow, err := client.Can(userID, action)
		if err != nil {
			log.Error(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if !allow {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		ctx.Next()
	}
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Cause   any    `json:"cause"`
}

func SentryErrorReporter() gin.HandlerFunc {
	return sentryErrorReporter()
}

func sentryErrorReporter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		hub := sentrygin.GetHubFromContext(c)
		detectedErrors := c.Errors

		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err.(*service.Error)
			if err.Cause != nil {
				hub.CaptureException(err.Cause)
			} else {
				hub.CaptureException(err)
			}
			return
		}
	}
}
