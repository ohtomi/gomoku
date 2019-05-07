package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{}

func (u *Upgrade) Execute(robotFactory *RobotFactory, response http.ResponseWriter, request *http.Request) error {
	if u.Protocol == "ws" {
		if err := exchangeWsUpgradeMessage(robotFactory, response, request); err != nil {
			return err
		}
	}

	return nil
}

func exchangeWsUpgradeMessage(robotFactory *RobotFactory, response http.ResponseWriter, request *http.Request) error {
	connection, err := wsUpgrader.Upgrade(response, request, nil)
	if err != nil {
		return err
	}
	robotFactory.Connection = connection

	return nil
}
