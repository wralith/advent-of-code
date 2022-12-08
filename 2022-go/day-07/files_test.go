package day07

import (
	"testing"
)

func Test_spaceOfFoldersSmallerThan(t *testing.T) {

	data := readInput()
	m := inputToMap(data)
	got := spaceOfFoldersSmallerThan(10000, m)
	want := 1778099

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
