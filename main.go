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
