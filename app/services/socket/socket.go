package socket

import (
	"fmt"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

type Ws struct {
	server socketio.Server
}

func NewWebsocketService() *Ws {
	server := socketio.NewServer(nil)
	ws := Ws{server: *server}
	ws.server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	go server.Serve()
	defer server.Close()
	http.Handle("/ws", server)
	return &ws
}
