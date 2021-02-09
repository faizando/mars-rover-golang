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

	// read file passed in arg
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	p := getPlateauCoordinates(strings.Split(string(input), "\n")[:1][0])
	// fmt.Println(p)

	rl := getListOfRovers(strings.Split(string(input), "\n")[1:])
	// fmt.Println(rl)

	processRoversAndPrintResult(p, rl)
	// fmt.Println(result)
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

func processRoversAndPrintResult(p plateau, roverlist []rover) {

	rl := roverlist

	for _, r := range rl {

		// TODO remove? rover checking from the array so that can simply loop on it and will not need: if cri != ri

		// crl := rl[:i-1]
		// crl = append(crl, rl[i:])

		// rover within plateau and >= 0,0
		if r.x >= 0 && r.x <= p.x && r.y >= 0 && r.y <= p.y {

			// haltCommands := false
			for _, com := range r.com {

				r.processCommand(string(com))

				// TODO when only one rover don't compare
				// TODO collision detection

				// when more than one rover compare for collisions
				// compare to cr (control rover)
				// for cri, cr := range crl {

				// index not same so that do not check collision with itself
				// TODO should check will not collide with the most up to date rover x,r not the control because here control is original copy of the input!
				// TODO will rl in forloop update if i change it within the loop here??

				// if cri != ri {

				// 	// TODO when there is only one rover in thelist will compare to itself
				// 	//TODO  && !roverWillGoOutOfBounds(r, p)

				// 	if !roverWillCollide(r, cr) {
				// 		r.processCommand(string(com))
				// 	} else {
				// 		haltCommands = true
				// 		break
				// 	}
				// }

				// }

				// if haltCommands == true {
				// 	break
				// }

			}

		}

		r.print()

	}

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

func roverWillCollide(r rover, cr rover) bool {
	r2 := r
	r2.move()
	if r2.x == cr.x && r2.y == cr.y {
		return true
	}
	return false
}

//TODO move *rover functions into its own file
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
