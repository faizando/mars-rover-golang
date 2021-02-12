package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type plateau struct {
	x, y int
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
