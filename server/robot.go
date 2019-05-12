package server

import (
	"regexp"
)

func (r *Robots) Run(conn UpgradedConnection) {
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		robot, found := r.SelectRobotItem(mt, string(message))
		if !found {
			continue
		}
		err = conn.WriteMessage(robot.Sink.Type, []byte (robot.Sink.Body))
		if err != nil {
			break
		}
	}
}

func (r *Robots) SelectRobotItem(messageType int, messageBody string) (*RobotItem, bool) {
	for _, element := range *r {
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
