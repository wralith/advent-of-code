package day03

import "testing"

func Test_findTotalRucksackPriorityFromFile(t *testing.T) {
	got, _ := findTotalRucksackPriorityFromFile("input.txt")
	want := 7967
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
