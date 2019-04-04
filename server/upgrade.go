package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func (u *Upgrade) Execute(robot *LiveRobot, response http.ResponseWriter, request *http.Request) error {
	if len(u.Scenario) == 0 {
		return nil
	}

	if err := u.exchangeUpgradeMessage(robot, response, request); err != nil {
		return err
	}

	return nil
}

func (u *Upgrade) exchangeUpgradeMessage(robot *LiveRobot, response http.ResponseWriter, request *http.Request) error {
	connection, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		return err
	}
	robot.Connection = connection

	return nil
}
