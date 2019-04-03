package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func (u *Upgrade) Accept(robot *Robot, response http.ResponseWriter, request *http.Request) error {
	if len(u.Scenario) == 0 {
		return nil
	}

	upgrader := websocket.Upgrader{}
	connection, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		return err
	}
	robot.Connection = connection

	return nil
}
