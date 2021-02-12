package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rover struct {
	x, y     int
	dir, com string
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

func getListOfRovers(input []string) []rover {
	rl := []rover{}

	// flatten 2 line input per rover into one line
	for i := 1; i < len(input); i += 2 {

		ri := input[i-1] + " " + input[i]
		rl = append(rl, getRover(ri))
	}

	return rl
}

func getRover(roverInput string) rover {
	validDirs := "NEWS"
	char := strings.Fields(roverInput)

	// Assumption: rover always given coordinates x,y, direction, and command(s)
	if len(char) != 4 {
		fmt.Println("Error, rover input values provided ", char, "are invalid")
		os.Exit(1)
	}

	r := rover{}

	if cx, err := strconv.Atoi(char[0]); err == nil {
		if cy, err := strconv.Atoi(char[1]); err == nil {

			// Assumption will allow rovers to land outside 0,plateau, life
			r.x = cx
			r.y = cy
			r.com = char[3]

			if strings.Contains(validDirs, char[2]) {
				r.dir = char[2]
			} else {
				fmt.Println("Error, rover input direction '", char[2], "' is invalid")
				os.Exit(1)
			}

		} else {
			fmt.Println("Error, could not read rover input values", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Error, could not read rover input values", err)
		os.Exit(1)
	}

	return r
}

func roverWillCollide(r rover, cr rover) bool {
	r2 := r
	r2.move()
	if r2.x == cr.x && r2.y == cr.y {
		return true
	}
	return false
}
