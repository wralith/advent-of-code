package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var mutex sync.Mutex

func main() {
	lastState := playWithNewMachine()

	fmt.Println(getTopCrates(lastState))
}

func getTopCrates(state map[int][]rune) (res string) {
	for i := 1; i <= len(state); i++ {
		res += string(state[i][len(state[i])-1])
	}
	return
}

func play() map[int][]rune {
	data, instructions := readDataAndInstructions()
	crates := parseHeader(data)
	for _, ins := range instructions {
		move, from, to := parseInstruction(ins)
		for i := 0; i < move; i++ {
			mutex.Lock()
			n := len(crates[from]) - 1
			crates[to] = append(crates[to], crates[from][n])
			crates[from] = crates[from][:n]
			mutex.Unlock()
		}
	}
	return crates
}

func parseHeader(raw []string) map[int][]rune {
	m := make(map[int][]rune)
	for _, line := range raw {
		where := 0
		for i, char := range line {
			if i%4 == 1 {
				where++
				if char != 32 {
					m[where] = append(m[where], char)
				}
			}
		}
	}
	return m
}

func parseInstruction(raw string) (move, from, to int) {
	fmt.Sscanf(raw, "move %d from %d to %d", &move, &from, &to)
	return
}

func readDataAndInstructions() ([]string, []string) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("unable to read file")
	}
	input := strings.Split(string(file), "\n\n")

	data := strings.Split(input[0], "\n")
	data = data[:len(data)-1]
	// reverse
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	instructions := strings.Split(input[1], "\n")

	return data, instructions
}

// PART TWO
func playWithNewMachine() map[int][]rune {
	data, instructions := readDataAndInstructions()
	crates := parseHeader(data)
	for _, ins := range instructions {
		move, from, to := parseInstruction(ins)

		mutex.Lock()
		n := len(crates[from])
		crates[to] = append(crates[to], crates[from][n-move:]...)
		crates[from] = crates[from][:n-move]
		mutex.Unlock()
	}
	return crates
}
