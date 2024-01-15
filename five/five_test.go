package main

import "testing"

func TestClosestLocation(t *testing.T) {
	want := 35
	got := closestLocation("almanac.test.txt")

	if want != got {
		t.Fatalf("wanted: %v, got: %v", want, got)
	}
}

func TestClosestLocaitonRanges(t *testing.T) {
	want := 46
	got := closestLocationRanges("almanac.test.txt")

	if want != got {
		t.Fatalf("wanted: %v, got: %v", want, got)
	}
}
