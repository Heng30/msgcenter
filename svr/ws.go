package svr

import (
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var tlock sync.Mutex
var topicMap = make(map[string][]*websocket.Conn)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ws(r gin.IRouter) {
	r.GET("/", func(c *gin.Context) {
		topic, ok := c.GetQuery("topic")
		if !ok || topic == "" {
			c.JSON(http.StatusBadRequest, errorBody(errors.New("can not found topic in url")))
			return
		}
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorBody(err))
			return
		}

		log.Printf("%s connect...", conn.RemoteAddr())
		tlock.Lock()
		if conns, ok := topicMap[topic]; ok {
			topicMap[topic] = append(conns, conn)
		} else {
			topicMap[topic] = []*websocket.Conn{conn}
		}
		tlock.Unlock()

		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				tlock.Lock()
				defer tlock.Unlock()
				removeConn(topic, conn)
				return
			}
		}
	})
}

func removeConn(topic string, conn *websocket.Conn) {
	if conns, ok := topicMap[topic]; ok {
		for i, item := range conns {
			if item == conn {
				topicMap[topic] = append(conns[:i], conns[i+1:]...)
				log.Printf("%s exit...", conn.RemoteAddr())
				conn.Close()
				return
			}
		}
	}
}

func broadcast(topic, msg string) {
	if topic == "" || msg == "" {
		return
	}

	tlock.Lock()
	defer tlock.Unlock()
	if conns, ok := topicMap[topic]; ok {
		errConns := []*websocket.Conn{}

		for _, conn := range conns {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				errConns = append(errConns, conn)
			}
		}

		for _, conn := range errConns {
			removeConn(topic, conn)
		}
	}
}
