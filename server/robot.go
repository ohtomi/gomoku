package server

import (
	"regexp"
)

func (r *Robots) Run(robotFactory *RobotFactory) {
	defer robotFactory.Connection.Close()

	for {
		mt, message, err := robotFactory.Connection.ReadMessage()
		if err != nil {
			break
		}
		robot, found := selectRobotItem(r, mt, string(message))
		if !found {
			continue
		}
		err = robotFactory.Connection.WriteMessage(robot.Sink.Type, []byte (robot.Sink.Body))
		if err != nil {
			break
		}
	}
}

func selectRobotItem(robots *Robots, messageType int, messageBody string) (*RobotItem, bool) {
	for _, element := range *robots {
		if element.Source.Type != messageType {
			continue
		}
		if len(element.Source.Body) != 0 {
			if !regexp.MustCompile(element.Source.Body).MatchString(messageBody) {
				continue
			}
		}
		return &element, true
	}
	return nil, false
}
