package main

import "testing"

var test = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"custom-results", 250.0, 50.0, 5.0, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range test {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one.")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one:", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("did not get expected value. got %f, expected %f", got, tt.expected)
		}
	}
}
