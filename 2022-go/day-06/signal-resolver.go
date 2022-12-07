package day06

import (
	"os"
)

func readFile(fileName string) string {
	file, _ := os.ReadFile(fileName)
	return string(file)
}

func findNonRepeated(data string, interval int) int {
	slice := []byte{}
	for i := 0; i < len(data); i++ {
		if i < interval {
			slice = append(slice, data[i])
		} else {
			slice = append(slice, data[i])
			slice = slice[1:]
			if !isContainsRepeated(slice) {
				return i + 1
			}
		}
	}
	return -1
}

func isContainsRepeated(slice []byte) bool {
	m := make(map[byte]bool)
	for _, v := range slice {
		m[v] = true
	}
	return len(m) < len(slice)
}

func resolveSignal() int {
	signal := readFile("input.txt")
	return findNonRepeated(signal, 4)
}

// PART TWO
func resolveMessage() int {
	signal := readFile("input.txt")
	return findNonRepeated(signal, 14)
}
