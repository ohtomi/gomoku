package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Robot struct {
	Connection *websocket.Conn
}

func RunRobots(upgrade *Upgrade, reporter Reporter, w http.ResponseWriter, r *http.Request) {
	robot := &Robot{}

	if upgrade != nil {
		if err := upgrade.Accept(robot, w, r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			reporter.Error(err.Error())
		}
	}

	if robot.Connection != nil {
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
