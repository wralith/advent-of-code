package day07

import (
	"os"
	"strconv"
	"strings"
)

func readInput() []string {
	data, _ := os.ReadFile("input.txt")
	return strings.Split(string(data), "\n")
}

// Answer Here!
func spaceOfFoldersSmallerThan(target int, data map[string]int) (result int) {
	for _, dir := range data {
		if dir <= 100000 {
			result += dir
		}
	}
	return
}

func inputToMap(data []string) map[string]int {
	path := []string{"/"}
	m := make(map[string]int)
	dirsCount := 0

	for i := 0; i < len(data); i++ {
		if data[i] == "$ cd .." {
			path = path[:len(path)-1] // back
			continue
		}

		cmd := strings.Split(data[i], " ")

		if cmd[0] == "$" {
			if cmd[1] == "cd" {
				path = move(cmd, m, path, i) // files in different dirs can have same name! => `filename+i`
			}
			continue
		}

		if cmd[0] != "dir" {
			addFileSize(cmd, m, path)
		}

		if cmd[0] == "dir" {
			dirsCount++
		}
	}

	return m
}

func addFileSize(cmd []string, dirs map[string]int, path []string) {
	size, _ := strconv.ParseInt(cmd[0], 10, 64)
	for _, p := range path {
		dirs[p] += int(size)
	}
}

func move(cmd []string, dirs map[string]int, path []string, count int) []string {
	new := cmd[2]
	if _, ok := dirs[new]; ok {
		new = cmd[2] + strconv.Itoa(count)
	}
	path = append(path, new)
	return path
}
