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
	x, y                int
	direction, commands string
}

func main() {

	// validDirections := "NEWS"
	// validCommands := "LRM"

	rl := []rover{}

	// read file passed in arg
	dat, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	l := strings.Split(string(dat), "\n")
	p := getPlateauCoordinates(l[:1][0]) // plateauLine := l[:1][0]
	fmt.Println(p)

	// save rover list
	rl = append(rl, rover{1, 2, "N", "LM"})
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
