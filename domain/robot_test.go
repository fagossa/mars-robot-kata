package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
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
