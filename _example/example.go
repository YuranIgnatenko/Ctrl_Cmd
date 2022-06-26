package main

import (
	. "ctrl_cmd"
	"time"
)

func main() {
	i := 0
	for {
		Line("clear")
		i++
		LineColor("hello", "red", "bg")
		LineColor("hello", "green", "bg")
		LineColor("hello", "blue", "bg")
		LineColor("hello", "yellow", "bg")
		LineColor("hello", "cyan", "bg")
		LineColor("hello", "black", "bg")
		LineColor("hello", "white", "bg")
		LineColor("hello", "violet", "bg")

		LineColor("hello", "red", "fg")
		LineColor("hello", "green", "fg")
		LineColor("hello", "blue", "fg")
		LineColor("hello", "yellow", "fg")
		LineColor("hello", "cyan", "fg")
		LineColor("hello", "black", "fg")
		LineColor("hello", "white", "fg")
		LineColor("hello", "violet", "fg")

		Line("echo `tput setab 1`Hello`tput sgr0`")
		Line("echo `tput setab 2`Hello`tput sgr0`")
		Line("echo `tput setaf 3``tput setab 5`Hello`tput sgr0`")

		Line(F("echo \"%v\"", Block))
		Line(F("echo \"%v\"", TriangleL))
		Line(F("echo \"%v\"", TriangleR))
		Line(F("echo \"%v\"", TriangleU))
		Line(F("echo \"%v\"", TriangleD))

		pid := LineGet("ps aux | grep main | grep -v color")

		Line(F("echo [isec:%v] -- [pid:%v]", i, pid))
		time.Sleep(time.Second * 1)
	}
}
