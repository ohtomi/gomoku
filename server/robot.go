package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type RobotFactory struct {
	Connection *websocket.Conn
}

func TryRunRobots(upgrade *Upgrade, reporter Reporter, w http.ResponseWriter, r *http.Request) {
	robotFactory := &RobotFactory{}

	if upgrade != nil {
		if err := upgrade.Execute(robotFactory, w, r); err != nil {
			reporter.Error(err.Error())
			return
		}

		if robotFactory.Connection == nil {
			reporter.Error("TODO")
			return
		}
		defer robotFactory.Connection.Close()

		for {
			mt, message, err := robotFactory.Connection.ReadMessage()
			if err != nil {
				break
			}
			err = robotFactory.Connection.WriteMessage(mt, message)
			if err != nil {
				break
			}
		}
	}
}
