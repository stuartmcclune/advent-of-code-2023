package main

import "testing"

func TestCalibrate(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	want := uint64(142)
	output := calibrate(input)

	if want != output {
		t.Fatalf(`Got %v, want %v`, output, want)
	}
}

func TestCalibrate2(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	want := uint64(281)
	output := calibrate2(input)

	if want != output {
		t.Fatalf(`Got %v, want %v`, output, want)
	}
}
