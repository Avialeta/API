package api

import (
	"golang.org/x/net/websocket"

	"github.com/avialeta/api/log"
)

func init() {
	mux.Handle("/statistic", websocket.Handler(handleStatistic))
}

func handleStatistic(ws *websocket.Conn) {
	for k, v := range ws.Config().Header {
		log.Debug.Println(k, " - ", v)
	}

	ws.Write([]byte("Hi"))
}
