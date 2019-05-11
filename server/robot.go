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
		defer robotFactory.Connection.Close()

		for {
			mt, message, err := robotFactory.Connection.ReadMessage()
			if err != nil {
				break
			}
			robot, found := robots.SelectRobotItem(mt, string(message))
			if !found {
				continue
			}
			err = robotFactory.Connection.WriteMessage(robot.Sink.Type, []byte (robot.Sink.Body))
			if err != nil {
				break
			}
		}
	}
}
