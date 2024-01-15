package main

import "testing"

func assertEqual[U comparable](want, got U, t *testing.T) {
	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func TestEleven(t *testing.T) {
	want := 374
	got := eleven("input.test.txt", 1)

	assertEqual(want, got, t)
}
