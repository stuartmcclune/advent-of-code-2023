package main

import "testing"

func TestTwoOne(t *testing.T) {
	fname := "input1.test.txt"
	want := 8

	output := twoOne(fname, counts{red: 12, green: 13, blue: 14})

	if output != want {
		t.Fatalf("Wanted: %v, got: %v", want, output)
	}
}

func TestTwoTwo(t *testing.T) {
	fname := "input1.test.txt"
	want := 2286

	output := twoTwo(fname)

	if output != want {
		t.Fatalf("Wanted: %v, got: %v", want, output)
	}
}
