package domain

import (
	"fmt"
)

// Robot represents the actual drone that is capable to move and rotate
// It supports the directions: N,E,W,S and the commands F,B,R,L
type Robot struct {
	Name string

	// TODO: add additional fields as per the requirements
}

// NewRobot allows to create valid robots. It receives the robot initial 
// state and the actions it must execute.
//
// conf, follows this format "X Y Dir". For example, 5 5 N
// commands, are the orders to execute. For example, FFRFFL
func NewRobot(conf, commands string) (*Robot, error) {
	if len(conf) == 0 {
		// fails because the configuration is invalid
		return nil, fmt.Errorf("robot conf is mandatory: <%s>", conf)
	}

	// TODO: complete the parsing for creating new robots

	return nil, nil
}

// RunCommands execute the list of commands in a synchronous way
func (r Robot) RunCommands() {
	// TODO: make the robot execute the slice of commands
	// for i, c := range r.Commands {
	// 	...
	// }
}

func (r Robot) Hello() string {
	return fmt.Sprintf("Hello, %s", r.Name)
}
