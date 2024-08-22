package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

type GreetService struct{}

func (g *GreetService) Greet(name string) string {
	// wait for 3 seconds
	time.Sleep(time.Second * 3)
	robotgo.WriteAll("Hello, World!")
	// press Command key
	robotgo.KeyDown("command")
	time.Sleep(10 * time.Millisecond)
	// press V key
	robotgo.KeyTap("v")
	time.Sleep(10 * time.Millisecond)
	// release Command key
	robotgo.KeyUp("command")
	return "Hello " + name + "!"
}
