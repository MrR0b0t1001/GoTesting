package main

import (
	"testing"
)

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"division-by-zero", 100.0, 0.0, -1.0, true},
	{"produces-fraction", -1.0, -2.0, 0.5, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Expected no error but got one")
			}
		}

		if got != tt.expected {
			t.Errorf("Expected: %f , Got: %f", tt.expected, got)
		}
	}
}
