package main

import (
	"testing"
	"fmt"
)

func TestRoll (t *testing.T) {
	var tests = []struct {
		i int
		e bool
	} {
		{ -1,		true },
		{ 0,		true },
		{ 1,		false },
		{ 100,		false },
		{ 8798237,	false },
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Test intput: %v", tt.i)
		t.Run(testname, func(t *testing.T) {
			out, err := roll(tt.i)
			if (err != nil) != tt.e {
				if(err != nil){
					t.Errorf("TEST FAILED: wanted an error, but received nil")
				} else {
					t.Errorf("TEST FAILED: did not want error, but received: %v", err)
				}
			}
			if out > tt.i && err == nil {
				t.Errorf("TEST FAILED: Rolled a max of %v, but got %v", tt.i, out)
			}
		})
	}
}

func TestMinRoll (t *testing.T) {
	var tests = []struct {
		i, min int
		e bool
	} {
		{ -1,		-2, 		true },
		{ -1,		1, 			true },
		{ 0,		0, 			true },
		{ 0, 		-1,			true },
		{ 1,		0,			true },
		{ 2,		1,			false },
		{ 100,		90, 		false },
		{ 8798237,	8798236, 	false },
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Test min intput: %v, %v", tt.i, tt.min)
		t.Run(testname, func(t *testing.T) {
			out, err := minRoll(tt.i, tt.min)
			if (err != nil) != tt.e {
				if(tt.e != false){
					t.Errorf("TEST FAILED: wanted an error, but received nil")
				} else {
					t.Errorf("TEST FAILED: did not want error, but received: %v", err)
				}
			}
			if out > tt.i && tt.e == false{
				t.Errorf("TEST FAILED: Rolled a max of %v, but got %v", tt.i, out)
			}
			if out < tt.min && err == nil{
				t.Errorf("TEST FAILED: Out(%v) is less than min(%v).", out, tt.min)
			}
		})
	}
}

func TestAddRoll (t *testing.T) {
	var tests = []struct {
		i, add int
		e bool
	} {
		{ -1,		-2, 		true },
		{ -1,		1, 			true },
		{ 0,		0, 			true },
		{ 1,		0,			true },
		{ 5,		10,			false },
		{ 100,		90, 		false },
		{ 8798237,	8798236, 	false },
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Test add intput: %v, %v", tt.i, tt.add)
		t.Run(testname, func(t *testing.T) {
			out, err := addRoll(tt.i, tt.add)
			if (err!=nil) != tt.e {
				if(err != nil){
					t.Errorf("TEST FAILED: wanted an error, but received nil or different error: %v", err)
				} else {
					t.Errorf("TEST FAILED: did not want error, but received: %v", err)
				}
			}
			if out <= tt.add && err == nil{
				t.Errorf("TEST FAILED: Added value %v, but somehow got %v", tt.add, out)
			}
		})
	}
}