package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func biggestCalorieFromFile(fileName string) (biggest int, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, fmt.Errorf("unable to open file `%s`", fileName)
	}

	return biggestCalorieSum(file), nil
}

func biggestCalorieSum(file *os.File) (biggest int) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	temp := 0
	for scanner.Scan() {
		if scanner.Text() != "" {
			num, _ := strconv.Atoi(scanner.Text())
			temp += num
			if temp > biggest {
				biggest = temp
			}
		} else {
			temp = 0
		}
	}

	return
}
