package main

import "testing"

func TestEightOne(t *testing.T) {
	want := 2
	got := eightOne("input.test.txt")

	if want != got {
		t.Fatalf("1: Wanted %v, got %v", want, got)
	}

	want = 6
	got = eightOne("input2.test.txt")

	if want != got {
		t.Fatalf("2: Wanted %v, got %v", want, got)
	}
}

func TestEightTwo(t *testing.T) {
	want := 6
	got := eightTwo("input3.test.txt")

	if want != got {
		t.Fatalf("1: Wanted %v, got %v", want, got)
	}
}
