package day01

import "testing"

func Test_biggestCalorieFromFile(t *testing.T) {
	got, _ := biggestCalorieFromFile("input.txt")
	want := 71502
	if got != want {
		t.Errorf("got %d - want %d", got, want)
	}

	got, _ = biggestCalorieFromFile("simple_input.txt")
	want = 22
	if got != want {
		t.Errorf("got %d - want %d", got, want)
	}

	_, err := biggestCalorieFromFile("none.txt")
	if err == nil {
		t.Errorf("got nil want error")
	}
}
