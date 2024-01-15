package main

import "testing"

func TestFifteenOne(t *testing.T) {
	want := 1320
	got := fifteenOne("input.test.txt")

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func TestFifteenTwo(t *testing.T) {
	want := 145
	got := fifteenTwo("input.test.txt")

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
