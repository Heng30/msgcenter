package svr

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	_ "log"
	"net/http"
)

func topic(r gin.IRouter) {
	r.POST("/topic/:topic", func(c *gin.Context) {
		topic := c.Param("topic")
		msg := make([]byte, 0, 4096)
		data := make([]byte, 1024)
		for {
			n, err := c.Request.Body.Read(data)
			if n > 0 {
				msg = append(msg, data[:n]...)
			}

			if err != nil {
				if err != io.EOF {
					c.JSON(http.StatusBadRequest, errorBody(err))
					return
				} else {
					break
				}
			}
			if n <= 0 {
				break
			}
		}

		if len(msg) <= 0 {
			c.JSON(http.StatusBadRequest, errorBody(errors.New("Not support empty body")))
			return
		}

		broadcast(topic, string(msg))
		c.JSON(http.StatusOK, gin.H{"code": 0})
	})
}
