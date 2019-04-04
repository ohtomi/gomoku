package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type LiveRobot struct {
	Connection *websocket.Conn
}

func RunRobots(upgrade *Upgrade, reporter Reporter, w http.ResponseWriter, r *http.Request) {
	robot := &LiveRobot{}

	if upgrade != nil {
		if err := upgrade.Execute(robot, w, r); err != nil {
			reporter.Error(err.Error())
			return
		}

		if robot.Connection == nil {
			reporter.Error("TODO")
			return
		}
		defer robot.Connection.Close()

		for {
			mt, message, err := robot.Connection.ReadMessage()
			if err != nil {
				break
			}
			err = robot.Connection.WriteMessage(mt, message)
			if err != nil {
				break
			}
		}
	}
}
