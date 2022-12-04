package day03

import (
	"bufio"
	"fmt"
	"os"
)

func findTotalBadgesPriorities(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return -1, fmt.Errorf("file named %s not found in current directory", fileName)
	}
	res := badgesPrioritySumFromInput(file)
	return res, nil
}

func badgesPrioritySumFromInput(file *os.File) (biggest int) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	i := 0
	temp := ""
	for scanner.Scan() {
		i++
		switch i % 3 {
		case 1:
			temp = scanner.Text()
		case 2:
			common := findCommons(temp, scanner.Text())
			temp = string(common)
		case 0:
			common := findCommons(temp, scanner.Text())
			temp = string(common)
			biggest += asciToPriorityPoint(rune(temp[0]))
		}
	}
	return
}

func findCommons(f, s string) (res string) {
	visited := make(map[rune]bool)
	for _, char := range f {
		visited[char] = true
	}
	for _, char := range s {
		if visited[char] {
			res += string(char)
		}
	}
	return
}
