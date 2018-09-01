package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"github.com/gin-gonic/gin"
)

func ApiServer() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "bitbot, baby!")
	})
	r.Run()
}

var APIServer = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && m.Content == "!api"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, "Starting API server...")
		go ApiServer()
		return false
	},
}
