package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestGetPlateauCoordinates(t *testing.T) {
	inputLine := "5 5"
	output, err := getPlateauCoordinates(inputLine)

	if err == nil && output != (plateau{x: 5, y: 5}) {
		t.Errorf("Expected plateau to be {5 5}, but got %v", output)
	}

	type testErrors struct {
		input       string
		expectedErr error
	}

	tests := []testErrors{
		{input: inputLine, expectedErr: nil},
		{input: "5", expectedErr: errors.New("Error, input plateau coordinates are invalid, expected `x y` where x and y represent integer value when x>=0 and y>=0")},
		{input: "5 5 5 5", expectedErr: errors.New("Error, input plateau coordinates are invalid, expected `x y` where x and y represent integer value when x>=0 and y>=0")},
		{input: "5 -5", expectedErr: errors.New("Error, input plateau coordinates are invalid, expected `x y` where x and y represent integer value when x>=0 and y>=0")},
		{input: "blah 5", expectedErr: errors.New("Error, input plateau x coordinate is invalid")},
		{input: "5 blah", expectedErr: errors.New("Error, input plateau y coordinate is invalid")},
	}

	for _, tc := range tests {
		_, gotErr := getPlateauCoordinates(tc.input)
		if !reflect.DeepEqual(gotErr, tc.expectedErr) {
			t.Errorf("Expected: %v, \n Got: %v", tc.expectedErr, gotErr)
		}
	}

}
