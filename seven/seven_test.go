package main

import "testing"

func TestSevenOne(t *testing.T) {
	want := 6440
	got := sevenOne("input.test.txt")

	if want != got {
		t.Fatalf("Wated %v, got %v", want, got)
	}
}

func TestSevenTwo(t *testing.T) {
	want := 5905
	got := sevenTwo("input.test.txt")

	if want != got {
		t.Fatalf("Wated %v, got %v", want, got)
	}
}
