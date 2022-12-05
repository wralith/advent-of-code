package main

import (
	"testing"
)

func Test_findContainingInFile(t *testing.T) {
	got := findContainingInFile("input.txt")
	want := 494
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func Test_findAnyOverlapInFile(t *testing.T) {
	got := findAnyOverlapInFile("input.txt")
	want := 833
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}
