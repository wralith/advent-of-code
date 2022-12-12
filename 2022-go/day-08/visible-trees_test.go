package day08

import "testing"

func Test_findVisibleTreeCount(t *testing.T) {
	got := findVisibleTreeCount()
	want := 1708

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
