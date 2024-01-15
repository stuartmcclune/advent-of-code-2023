package main

import "testing"

func TestSumPartNumbers(t *testing.T) {
	input := "input1.test.txt"
	want := 4361
	got := sumPartNumbers(input)

	if got != want {
		t.Fatalf("Wanted: %v, got: %v", want, got)
	}
}

func TestIsDigit(t *testing.T) {
	if isDigit(".") {
		t.Fatalf("'.' is not a digit")
	}
	if isDigit("*") {
		t.Fatalf("'*' is not a digit")
	}
	if !isDigit("9") {
		t.Fatalf("'9' is a digit")
	}
}

func TestIsSymbol(t *testing.T) {
	if isSymbol(".") {
		t.Fatalf("'.' is not a symbol")
	}
	if isSymbol("9") {
		t.Fatalf("'9' is not a symbol")
	}
	if !isSymbol("+") {
		t.Fatalf("'+' is a symbol")
	}
}

func TestSumGearRatios(t *testing.T) {
	want := 467835
	got := sumGearRatios("input1.test.txt")
	if want != got {
		t.Fatalf("Want: %v, got: %v", want, got)
	}
}
