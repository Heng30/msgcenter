package middlewares

import (
	"github.com/gin-gonic/gin"
	"local/db"
	"net/http"
)

func Auth(testMode bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method != "GET" && method != "POST" {
			c.AbortWithStatus(http.StatusMethodNotAllowed)
		}

		var token string
		if up := c.Request.Header.Get("Upgrade"); up == "websocket" {
            token, _ = c.GetQuery("token")
		} else {
			token = c.Request.Header.Get("Auth-Token")

		}

		if token == "" {
			c.AbortWithStatus(http.StatusNonAuthoritativeInfo)
		}
		if !canAccess(token, testMode) {
			c.AbortWithStatus(http.StatusNotFound)
		}

		c.Next()
	}
}

func canAccess(token string, testMode bool) bool {
	if testMode {
		return token == "testToken"
	}

	if token == "testToken" {
		return false
	}

	if tokens, err := db.QueryAccessTokens(); err == nil {
		for _, item := range tokens {
			if item == token {
				return true
			}
		}
	}
	return false
}
