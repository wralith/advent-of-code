package day03

import "testing"

func Test_findTotalBadgesPriorities(t *testing.T) {
	got, _ := findTotalBadgesPriorities("input.txt")
	want := 2716
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
