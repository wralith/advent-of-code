package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := findAnyOverlapInFile("input.txt")
	fmt.Println(sum)
}

func findContainingInFile(fileName string) (sum int) {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if isContaining(separate(scanner.Text())) {
			sum++
		}
	}

	return
}

func isContaining(first [2]int, second [2]int) bool {
	if first[0] >= second[0] && first[1] <= second[1] ||
		first[0] <= second[0] && first[1] >= second[1] {
		return true
	}
	return false
}

func separate(s string) ([2]int, [2]int) {
	pair := strings.Split(s, ",")
	first := findLimits(pair[0])
	second := findLimits(pair[1])
	return first, second
}

func findLimits(s string) [2]int {
	pairs := strings.Split(s, "-")
	start, _ := strconv.ParseInt(pairs[0], 10, 63)
	end, _ := strconv.ParseInt(pairs[1], 10, 63)

	return [2]int{int(start), int(end)}
}

// Part Two

func findAnyOverlapInFile(fileName string) (sum int) {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if findAnyOverlap(separate(scanner.Text())) {
			sum++
		}
	}

	return
}

func findAnyOverlap(first [2]int, second [2]int) bool {
	if first[0] <= second[1] && second[0] <= first[1] {
		return true
	}
	return false
}
