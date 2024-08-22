// Code generated automatically. DO NOT EDIT.
package disableiface

import (
	"time"
)

type Greeter interface {
	SayHi() string
}

type CostcoGreeter struct {
	Year int
}

func (c *CostcoGreeter) SayHi() string {
	greeting := "Welcome to Costco"

	if c.Year >= 2505 {
		greeting += ", I love you"
	}

	return greeting + "."
}

var (
	greeter       Greeter
	otherGreeter  Greeter
	costcoGreeter *CostcoGreeter
	greeters      = []Greeter{greeter}
	boolChan      = make(chan bool)
	now           = time.Now().Unix()
)
