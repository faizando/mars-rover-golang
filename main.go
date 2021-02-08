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

	// validdirs := "NEWS"
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
	char := strings.Fields(roverInput)
	//todo if rover has no commands to save

	r := rover{}

	if cx, err := strconv.Atoi(char[0]); err == nil {
		if cy, err := strconv.Atoi(char[1]); err == nil {
			// Will throw error if rover lands in negative bounds
			if cx >= 0 && cy >= 0 {
				r.x = cx
				r.y = cy
				r.dir = char[2]
				r.com = char[3]
			} else {
				fmt.Println("Error, rover input value invalid")
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
