package main

import "testing"

func TestSixOne(t *testing.T) {
	want := 288
	got := sixOne("input.test.txt")

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func TestSixTwo(t *testing.T) {
	want := 71503
	got := sixTwo("input.test.txt")

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
