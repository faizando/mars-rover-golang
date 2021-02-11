package main

import (
	"testing"
)

func TestNewRover(t *testing.T) {
	r := rover{x: 1, y: 1, dir: "N", com: "LLM"}

	if r.x != 1 || r.y != 1 || r.dir != "N" || r.com != "LLM" {
		t.Errorf("Expected rover to be {x: 1, y: 1, dir: 'N', com: 'LLM'}, but got %v", r)
	}

}

func TestRoverSpin(t *testing.T) {

	r := rover{x: 1, y: 1, dir: "N", com: "LLM"}

	antiClockWise := "WSEN"
	for _, v := range antiClockWise {
		r.spin("L")
		if r.dir != string(v) {
			t.Errorf("Expected rover to be facing %v, but got %v", string(v), r.dir)
		}
	}

	clockWise := "ESWN"
	for _, v := range clockWise {
		r.spin("R")
		if r.dir != string(v) {
			t.Errorf("Expected rover to be facing %v, but got %v", string(v), r.dir)
		}
	}

}

func TestRoverMove(t *testing.T) {
	r := rover{x: 1, y: 1, dir: "N", com: "LLM"}

	r.move()
	if r.x != 1 || r.y != 2 {
		t.Errorf("Expected rover to be {x: 1, y: 2}, but got rover: %v", r)
	}

	r.dir = "S"
	r.move()
	if r.x != 1 || r.y != 1 {
		t.Errorf("Expected rover to be {x: 1, y: 1}, but got rover: %v", r)
	}

	r.dir = "E"
	r.move()
	if r.x != 2 || r.y != 1 {
		t.Errorf("Expected rover to be {x: 2, y: 1}, but got rover: %v", r)
	}

	r.dir = "W"
	r.move()
	if r.x != 1 || r.y != 1 {
		t.Errorf("Expected rover to be {x: 1, y: 1}, but got rover: %v", r)
	}

}
