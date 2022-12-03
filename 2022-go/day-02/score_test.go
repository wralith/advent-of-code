package day02

import "testing"

func Test_scanGame(t *testing.T) {
	got, _ := scanGame("input.txt")
	want := 10624
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
