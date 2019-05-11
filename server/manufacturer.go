package server

import (
	"regexp"
)

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
