package main

import "testing"

func assertEqual[U comparable](want, got U, t *testing.T) {
	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func TestArrangementsOne(t *testing.T) {
	assertEqual(1, parseRow("???.### 1,1,3").arrangements(), t)
}

func TestArrangementsTwo(t *testing.T) {
	assertEqual(4, parseRow(".??..??...?##. 1,1,3").arrangements(), t)
}

func TestArrangementsThree(t *testing.T) {
	assertEqual(1, parseRow("?#?#?#?#?#?#?#? 1,3,1,6").arrangements(), t)
}

func TestArrangementsFour(t *testing.T) {
	assertEqual(1, parseRow("????.#...#... 4,1,1").arrangements(), t)
}

func TestArrangementsFive(t *testing.T) {
	assertEqual(4, parseRow("????.######..#####. 1,6,5").arrangements(), t)
}

func TestArrangementsSix(t *testing.T) {
	assertEqual(10, parseRow("?###???????? 3,2,1").arrangements(), t)
}

func TestUnfoldArrangementsOne(t *testing.T) {
	r := parseRow("???.### 1,1,3")
	r = *r.unfold()
	assertEqual(1, r.arrangements(), t)
}

func TestUnfoldArrangementsTwo(t *testing.T) {
	r := parseRow("?###???????? 3,2,1")
	r = *r.unfold()
	assertEqual(506250, r.arrangements(), t)
}
