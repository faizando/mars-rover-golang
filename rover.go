package main

import (
	"fmt"
	"os"
	"strings"
)

type rover struct {
	x, y     int
	dir, com string
}

func (r *rover) processCommand(c string) {

	validCommands := "LRM"

	if strings.Contains(validCommands, c) {
		if c == "M" {
			r.move()
		} else {
			r.spin(c)
		}
	} else {
		fmt.Println("Error, command ", c, " to rover is invalid")
		os.Exit(1)
	}
}

func (r *rover) print() {
	fmt.Println(r.x, r.y, r.dir)
}

func (r *rover) spin(d string) {
	if d == "R" {
		switch r.dir {
		case "E":
			r.dir = "S"
		case "S":
			r.dir = "W"
		case "W":
			r.dir = "N"
		case "N":
			r.dir = "E"
		}
	} else if d == "L" {
		switch r.dir {
		case "S":
			r.dir = "E"
		case "W":
			r.dir = "S"
		case "N":
			r.dir = "W"
		case "E":
			r.dir = "N"
		}
	}
}

func (r *rover) move() {
	switch r.dir {
	case "E":
		r.x = r.x + 1
	case "W":
		r.x = r.x - 1
	case "N":
		r.y = r.y + 1
	case "S":
		r.y = r.y - 1
	}
}
