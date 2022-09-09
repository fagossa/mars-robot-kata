package domain

import (
	"fmt"
)

type Robot struct {
	Name string
}

func (r Robot) Hello() string {
	return fmt.Sprintf("Hello, %s", r.Name)
}
