package main

import (
	"coursecontent/demo/pkg/display"
	"coursecontent/demo/pkg/msg"
)

func main() {

	msg.Hi()
	display.Display("helo from display")
	msg.Exciting("an exciting message")

}
