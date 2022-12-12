package day08

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findVisibleTreeCount() int {
	data := readData()
	return visibleTrees(data)
}

func readData() []string {
	data, _ := os.ReadFile("input.txt")
	return strings.Fields(string(data))
}

func visibleTrees(cols []string) int {
	visible := make(map[string]bool)

	// vertical, starts from one since 0 is always visible
	for col := 1; col < len(cols)-1; col++ {
		tallest := byteToInt(cols[0][col])
		for row := 1; row < len(cols)-1; row++ {
			tallest = findTallest(visible, cols[row][col], tallest, row, col)
		}
		tallest = byteToInt(cols[len(cols)-1][col])
		for row := len(cols) - 2; row > 0; row-- {
			tallest = findTallest(visible, cols[row][col], tallest, row, col)
		}
	}
	// horizontal
	for row := 1; row < len(cols)-1; row++ {
		tallest := byteToInt(cols[row][0])
		for col := 1; col < len(cols[row])-1; col++ {
			tallest = findTallest(visible, cols[row][col], tallest, row, col)
		}
		tallest = byteToInt(cols[row][len(cols)-1])
		for col := len(cols[row]) - 2; col > 0; col-- {
			tallest = findTallest(visible, cols[row][col], tallest, row, col)
		}
	}

	return (len(cols) * 2) + (len(cols[0]) * 2) - 4 + len(visible)
}

func findTallest(m map[string]bool, cur byte, tall int, row, col int) int {
	s := byteToInt(cur)
	if s > tall {
		tall = s
		m[fmt.Sprintf("%d-%d", row, col)] = true
	}
	return tall
}

func byteToInt(b byte) int {
	i, _ := strconv.Atoi(string(b))
	return i
}
