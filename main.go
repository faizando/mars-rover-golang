package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type plateau struct {
	x, y int
}

type rover struct {
	x, y     int
	dir, com string
}

func main() {

	// validCommands := "LRM"

	// read file passed in arg
	input, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// l := strings.Split(string(input), "\n")[:1][0]
	p := getPlateauCoordinates(strings.Split(string(input), "\n")[:1][0]) // plateauLine := l[:1][0]
	fmt.Println(p)

	// save rover list
	rl := getListOfRovers(strings.Split(string(input), "\n")[1:])
	fmt.Println(rl)

	// process rover commands
	// todo
}

func getPlateauCoordinates(line string) plateau {
	char := strings.Fields(line)
	p := plateau{}

	if cx, err := strconv.Atoi(char[0]); err == nil {
		if cy, err := strconv.Atoi(char[1]); err == nil {
			if cx >= 0 && cy >= 0 {
				p.x = cx
				p.y = cy
			} else {
				fmt.Println("Error, input plateau coordinate invalid")
				os.Exit(1)
			}
		} else {
			fmt.Println("Error, could not read input plateau coordinate", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Error, could not read input plateau coordinate", err)
		os.Exit(1)
	}

	return p
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

func processRoversAndPrintResult(p plateau, rl []rover) {

	// move rover and when cannot move further and wont collide print it
	// so not looping twice, more elegant solution

	for i, r := range rl {

		// rover within plateau and >= 0,0
		if r.x >= 0 && r.x <= p.x && r.y >= 0 && r.y <= p.y {

			// compare to cr (control rover)
			for ci, cr := range rl {
				// index not same so that do not check collision with itself
				// TODO should check will not collide with the most up to date rover x,r not the control because here control is original copy of the input!
				if ci != i && roverWillNotCollide(r, cr) {
					processCommand(r)
					r.move()
					r.print()
				}
			}

		}
	}

}

func processCommand(r rover, c string) (rover, error) {
	validCommands := "LRM"

	rt := r

	if strings.Contains(validCommands, c) {
		switch c {
		case "L":
		case "R":
			rt.spin(c)
		case "M":
			// todo check collision here before moving?
			rt.move()
		}
	} else {
		fmt.Println("Error, command ", c, " to rover is invalid")
		os.Exit(1)
		// return rt, nil
	}

	return rt, nil

}

func roverWillNotCollide(r rover, cr rover) bool {
	r2 := r
	r2.move()
	if r2.x != cr.x && r2.y != cr.y {
		return true
	}
	return false
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
