package main

import "testing"

func TestNineOne(t *testing.T) {
	want := 114
	got := nineOne("input.test.txt")

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func TestNineTwo(t *testing.T) {
	want := 2
	got := nineTwo("input.test.txt")

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
