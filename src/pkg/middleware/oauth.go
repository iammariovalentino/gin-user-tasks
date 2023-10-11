package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/server"
)

type (
	// ErrorHandleFunc error handling function
	ErrorHandleFunc func(*gin.Context, error)
	// Config defines the config for Session middleware
	Config struct {
		// error handling when starting the session
		ErrorHandleFunc ErrorHandleFunc
		// keys stored in the context
		TokenKey string
		// defines a function to skip middleware.Returning true skips processing
		// the middleware.
		Skipper func(*gin.Context) bool

		// Oauth Server
		OauthServer *server.Server
	}
)

// HandleTokenVerify Verify the access token of the middleware
func HandleTokenVerify(config ...Config) gin.HandlerFunc {
	cfg := config[0]

	if cfg.ErrorHandleFunc == nil {
		cfg.ErrorHandleFunc = cfg.ErrorHandleFunc
	}

	tokenKey := cfg.TokenKey
	if tokenKey == "" {
		tokenKey = cfg.TokenKey
	}

	return func(c *gin.Context) {
		if cfg.Skipper != nil && cfg.Skipper(c) {
			c.Next()
			return
		}
		ti, err := cfg.OauthServer.ValidationBearerToken(c.Request)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
		}

		c.Set(tokenKey, ti)
		c.Next()
	}
}
