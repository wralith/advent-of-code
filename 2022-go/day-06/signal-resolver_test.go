package day06

import "testing"

func Test_resolveSignal(t *testing.T) {
	got := resolveSignal()
	want := 1655
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func Test_resolveMessage(t *testing.T) {
	got := resolveMessage()
	want := 2665
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
