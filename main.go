package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	// read file passed in arg
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	p := getPlateauCoordinates(strings.Split(string(input), "\n")[:1][0])
	rl := getListOfRovers(strings.Split(string(input), "\n")[1:])
	processRoversAndPrintResult(p, rl)

}

func processRoversAndPrintResult(p plateau, roverlist []rover) {

	rl := roverlist

	for i := range rl {

		willCol := false
		if rl[i].x >= 0 && rl[i].x <= p.x && rl[i].y >= 0 && rl[i].y <= p.y {

			for _, com := range rl[i].com {

				if !strings.Contains("LRM", string(com)) {
					fmt.Println("Error, command ", string(com), " to rover is invalid")
					os.Exit(1)
				}

				if string(com) != "M" {

					rl[i].spin(string(com))

				} else {

					for j := range rl {

						if i != j {
							if roverWillCollide(rl[i], rl[j]) {
								willCol = true
								break
							} else {
								willCol = false
							}
						}

					}

					if willCol {
						break
					} else {
						rl[i].move()
					}
				}

			}

		}

		rl[i].print()
	}

}
