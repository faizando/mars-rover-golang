package main

import (
	"errors"
	"strconv"
	"strings"
)

type plateau struct {
	x, y int
}

func getPlateauCoordinates(inputLine string) (p plateau, e error) {
	char := strings.Fields(inputLine)

	if len(char) != 2 {
		return p, errors.New("Error, input plateau coordinates are invalid, expected `x y` where x and y represent integer value when x>=0 and y>=0")
	}

	if cx, err := strconv.Atoi(char[0]); err == nil {
		if cy, err := strconv.Atoi(char[1]); err == nil {
			if cx >= 0 && cy >= 0 {
				p.x = cx
				p.y = cy
			} else {
				return p, errors.New("Error, input plateau coordinates are invalid, expected `x y` where x and y represent integer value when x>=0 and y>=0")
			}
		} else {
			return p, errors.New("Error, input plateau y coordinate is invalid")
		}
	} else {
		return p, errors.New("Error, input plateau x coordinate is invalid")

	}

	return p, nil
}
