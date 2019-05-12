package server

import (
	"net/http"
)

type RobotFactory struct {
	Connection UpgradedConnection
}

type UpgradedConnection interface {
	ReadMessage() (int, []byte, error)
	WriteMessage(int, []byte) error
	Close() error
}

func TryRunRobots(upgrade *Upgrade, robots *Robots, reporter Reporter, w http.ResponseWriter, r *http.Request) {
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

		go func(conn UpgradedConnection, robots *Robots) {
			defer conn.Close()

			for {
				mt, message, err := conn.ReadMessage()
				if err != nil {
					break
				}
				robot, found := robots.SelectRobotItem(mt, string(message))
				if !found {
					continue
				}
				err = conn.WriteMessage(robot.Sink.Type, []byte (robot.Sink.Body))
				if err != nil {
					break
				}
			}
		}(robotFactory.Connection, robots)
	}
}
