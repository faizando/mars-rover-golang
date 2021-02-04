package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type plateau struct {
	x int
	y int
}

func main() {

	dat, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	l := strings.Split(string(dat), "\n")
	plateauLine := l[:1][0]
	fmt.Println(plateauLine)
	fmt.Println(getPlateauCoordinates(plateauLine))
}

func getPlateauCoordinates(line string) (x int, y int) {
	char := strings.Fields(line)

	if s, err := strconv.Atoi(char[0]); err == nil {
		x = s
	} else {
		os.Exit(1)
	}

	if s, err := strconv.Atoi(char[1]); err == nil {
		y = s
	} else {
		os.Exit(1)
	}

	return x, y
}
