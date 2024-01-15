package main

import "testing"

func TestCalculatePoints(t *testing.T) {
	want := 13
	got := calculatePoints("input.test.txt")

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func TestCalculateNumberOfCards(t *testing.T) {
	want := 30
	got := calclateNumberOfCards("input.test.txt")

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
