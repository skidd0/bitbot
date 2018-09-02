package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"github.com/bbriggs/bitbot/bitbot/apiserver"
	"net/http"
	log "gopkg.in/inconshreveable/log15.v2"
)

var APIServer = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && m.Content == "!api"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, "API  server started...")
		hub := apiserver.NewHub(irc)
		go hub.Run()
		http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
			apiserver.ServeWs(hub, w, r)
		})
		err := http.ListenAndServe("localhost:8080", nil)
		if err != nil {
			log.Error(err.Error())
		}
		return true
	},
}
