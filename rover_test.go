package main

import (
	"fmt"
	"reflect"
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

func TestRoverWillCollide(t *testing.T) {
	r1 := rover{x: 1, y: 1, dir: "N", com: "LLM"}
	r2 := rover{x: 1, y: 2, dir: "N", com: "LLM"}
	r3 := rover{x: 3, y: 3, dir: "N", com: "LLM"}

	if !roverWillCollide(r1, r2) {
		t.Errorf("Expected rover to collide, got false")
	}

	if roverWillCollide(r1, r3) {
		t.Errorf("Expected rover to collide, got false")
	}
}

func TestGetRover(t *testing.T) {

	type testCase struct {
		input       string
		expectedRov rover
		expectedErr error
	}

	tr := rover{}

	tests := []testCase{
		{input: "1 1 N L", expectedRov: rover{x: 1, y: 1, dir: "N", com: "L"}, expectedErr: nil},
		{input: "1 1 N", expectedRov: tr, expectedErr: fmt.Errorf("Error, rover input values provided [1 1 N] are invalid")},
		{input: "1 1 blah L", expectedRov: tr, expectedErr: fmt.Errorf("Error, rover input direction 'blah' is invalid")},
		{input: "a 1 N L", expectedRov: tr, expectedErr: fmt.Errorf("Error, input rover x coordinate value of 'a' is invalid")},
		{input: "1 b N L", expectedRov: tr, expectedErr: fmt.Errorf("Error, input rover y coordinate value of 'b' is invalid")},
	}

	for _, tc := range tests {
		gotRov, gotErr := getRover(tc.input)
		if gotErr == nil && !reflect.DeepEqual(gotRov, tc.expectedRov) {
			t.Errorf("Expected: %v, \n Got: %v", tc.expectedRov, gotRov)
		}

		if !reflect.DeepEqual(gotErr, tc.expectedErr) {
			t.Errorf("Expected: %v, \n Got: %v", tc.expectedErr, gotErr)
		}
	}
}

func TestGetListOfRovers(t *testing.T) {

	validInput := []string{"1 2 N", "L", "3 3 E", "M"}
	expected := []rover{rover{x: 1, y: 2, dir: "N", com: "L"},
		rover{x: 3, y: 3, dir: "E", com: "M"}}

	got, _ := getListOfRovers(validInput)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %v, \n Got: %v", expected, got)
	}

	invalidInput := []string{"blah 2 N", "L", "3 3 E", "M"}
	expectedErr := fmt.Errorf("Error, input rover x coordinate value of 'blah' is invalid")
	_, err := getListOfRovers(invalidInput)
	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("Expected: %v, \n Got: %v", expectedErr, err)
	}
}
