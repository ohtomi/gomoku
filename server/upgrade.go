package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func (u *Upgrade) Accept(response http.ResponseWriter, request *http.Request) error {
	if len(u.Scenario) == 0 {
		return nil
	}

	upgrader := websocket.Upgrader{}
	connection, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		return err
	}
	defer connection.Close()

	for {
		mt, message, err := connection.ReadMessage()
		if err != nil {
			break
		}
		err = connection.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}

	return nil
}
