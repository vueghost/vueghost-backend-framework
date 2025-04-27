package System

import "fmt"

type Debug struct {
}

func (d *Debug) Error(title string, message string, err error) {
	_ = fmt.Errorf("Error %s %s %v", title, message, err)
}
