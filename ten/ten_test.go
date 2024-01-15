package main

import "testing"

func TestTenOne(t *testing.T) {
	want := 4
	got := tenOne("input.test.txt")

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	want = 8
	got = tenOne("input2.test.txt")

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func TestTenTwo(t *testing.T) {
	want := 4
	got := tenTwo("input3.test.txt")

	if want != got {
		t.Fatalf("3 Wanted %v, got %v", want, got)
	}

	want = 4
	got = tenTwo("input4.test.txt")

	if want != got {
		t.Fatalf("4 Wanted %v, got %v", want, got)
	}

	want = 8
	got = tenTwo("input5.test.txt")

	if want != got {
		t.Fatalf("5 Wanted %v, got %v", want, got)
	}

	want = 10
	got = tenTwo("input6.test.txt")

	if want != got {
		t.Fatalf("6 Wanted %v, got %v", want, got)
	}
}
