package day05

import (
	"testing"
)

func Test_play(t *testing.T) {
	game := play()
	got := getTopCrates(game)
	want := "CNSZFDVLJ"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func Test_playWithNewMachine(t *testing.T) {
	game := playWithNewMachine()
	got := getTopCrates(game)
	want := "QNDWLMGNS"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
