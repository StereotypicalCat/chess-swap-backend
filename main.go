package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/:name/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		// Send message to all participants on url other than me.
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			return (q.Request.URL.Path == s.Request.URL.Path) && (s != q)
		})
	})

	r.Run(":5000")
}
