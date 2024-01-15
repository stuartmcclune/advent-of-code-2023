package main

import "testing"

func TestFourteenOne(t *testing.T) {
	want := 136
	got := fourteenOne("input.test.txt")

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func TestFourteenTwo(t *testing.T) {
	want := 64
	got := fourteenTwo("input.test.txt")

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
