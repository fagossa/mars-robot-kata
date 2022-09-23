package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRobotString(t *testing.T) {
	testCases := []struct {
		name string
		robotName string
		expectedMsg string
	} {
		{
			name: "happy path",
			robotName: "Marie",
			expectedMsg: "Hello, Marie",
		},
	}
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			r := Robot { Name: c.robotName }
			assert.Equal(t, c.expectedMsg, r.Hello())
		})
	}
}


func TestNewRobot(t *testing.T) {
	testCases := []struct {
		name string
		robotConf string
		commands string
		expectedRobot *Robot
		isErr bool
	} {
		// TODO: you need to modify the implementation to make 
		// this test case work
		{
			name: "happy path",
			robotConf: "0 0 N",
			commands: "RRRR",
			isErr: false,
			// TODO: specify the robot with the expected values
			expectedRobot: &Robot{},
		},
		{
			name: "fail on empty conf",
			robotConf: "",
			commands: "",
			isErr: true,
			expectedRobot: nil,
		},
		{
			name: "fail on invalid conf",
			robotConf: "6 6 X", // direction X is not valid
			commands: "",
			isErr: true,
			expectedRobot: nil,
		},
	}
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			r, err := NewRobot(c.robotConf, c.commands)
			if c.isErr {
				require.Error(t, err)
			} else {
				assert.Equal(t, c.expectedRobot, r)
			}
		})
	}
}

