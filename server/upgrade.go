package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func (u *Upgrade) Execute(robotFactory *RobotFactory, response http.ResponseWriter, request *http.Request) error {
	if err := u.exchangeUpgradeMessage(robotFactory, response, request); err != nil {
		return err
	}

	return nil
}

func (u *Upgrade) exchangeUpgradeMessage(robotFactory *RobotFactory, response http.ResponseWriter, request *http.Request) error {
	connection, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		return err
	}
	robotFactory.Connection = connection

	return nil
}
