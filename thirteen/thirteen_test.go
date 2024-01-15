package main

import "testing"

func TestThirteenOne(t *testing.T) {
	want := 405
	got := thirteenOne("input.test.txt")

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func TestThirteenTwo(t *testing.T) {
	want := 400
	got := thirteenTwo("input.test.txt")

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
