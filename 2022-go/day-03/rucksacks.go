package day03

import (
	"bufio"
	"fmt"
	"os"
)

func findTotalRucksackPriorityFromFile(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return -1, fmt.Errorf("file named %s not found in current directory", fileName)
	}
	res := prioritySumFromInput(file)
	return res, nil
}

func prioritySumFromInput(file *os.File) (biggest int) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		pair := cutString(scanner.Text())
		common := findCommon(pair[0], pair[1])
		biggest += asciToPriorityPoint(common)
	}
	return
}

func asciToPriorityPoint(r rune) int {
	if r >= 97 && r <= 122 {
		return int(r - 96)
	}
	if r >= 65 && r <= 90 {
		return int(r - 38)
	}
	return 0
}

func cutString(s string) (res [2]string) {
	mid := len(s) / 2
	res[0] = s[:mid]
	res[1] = s[mid:]
	return
}

func findCommon(f, s string) rune {
	visited := make(map[rune]bool)
	for _, char := range f {
		visited[char] = true
	}
	for _, char := range s {
		if visited[char] {
			return char
		}
	}
	return -1
}
