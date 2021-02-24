# Mars Rover Exercise

A squad of robotic rovers are to be landed by NASA on a plateau on Mars.

This plateau, which is curiously rectangular, must be navigated by the rovers so that their on board cameras can get a complete view of the surrounding terrain to send back to Earth.

A rover's position is represented by a combination of an x and y co-ordinates and a letter representing one of the four cardinal compass points. The plateau is divided up into a grid to simplify navigation. An example position might be 0, 0, N, which means the rover is in the bottom left corner and facing North.

In order to control a rover, NASA sends a simple string of letters. The possible letters are 'L', 'R' and 'M'. 'L' and 'R' makes the rover spin 90 degrees left or right respectively, without moving from its current spot.

'M' means move forward one grid point, and maintain the same heading.

Assume that the square directly North from (x, y) is (x, y+1).

### Input:

The first line of input is the upper-right coordinates of the plateau, the lower-left coordinates are assumed to be 0,0.

The rest of the input is information pertaining to the rovers that have been deployed. Each rover has two lines of input. The first line gives the rover's position, and the second line is a series of instructions telling the rover how to explore the plateau.

The position is made up of two integers and a letter separated by spaces, corresponding to the x and y co-ordinates and the rover's orientation.

Each rover will be finished sequentially, which means that the second rover won't start to move until the first one has finished moving.

### Output:

The output for each rover should be its final co-ordinates and heading.

### Test Input file:

```
5 5

1 2 N

LMLMLMLMM

3 3 E

MMRMMRMRRM
```

### Expected Output to terminal:

```
1 3 N

5 1 E
```

This challenge extract is from [mars rover code archives](https://code.google.com/archive/p/marsrovertechchallenge/).

## Getting Started

Running it on your local machine

### Prerequisites

Have Go installed on machine and setup.

```
go version
```

This will print out your Go version. If you do not have it then see this guide to install the latest version: [Download and install Go](https://golang.org/doc/install)

### Building and running the program

In the project directory, run the following command to build the application into an executable file `mars-rover-golang`:

```
go build
```

The program takes a file as an argument which should contain the input. Run the program with input.txt file as input argument with following command:

```
./mars-rover-golang input.txt
```

This reads input from input.txt and prints output to the console.

To run a longer input data set provided in more-input.txt file run:

```
./mars-rover-golang more-input.txt
```

## Running Unit tests

Run unit tests by running

```
go test
```

## Assumptions and program features

- plateau's lower bounds are 0,0, and upper bounds are from first line of the input file
- first line of input file only contains the plateau coordinates
- following the first line, the pair of lines in the input file are:
  - a rover's starting coordinates
  - string of instructions for this rover in order
- program may fail if given blank lines or incorrect set of data.
- if rover lands out of bounds, then it will not move because it was not mentioned in the brief, and there may be another program to handle such scenarios to navigate the terrain if landed out of the plateau
- if moving a rover will result in a collision with another rover it will halt the program

## Future work

- could workout which coordinates of the plateau were unvisited by rovers in the initial run, and
- suggest path for rovers to cover this unvisited area (ideally with least number of moves)
