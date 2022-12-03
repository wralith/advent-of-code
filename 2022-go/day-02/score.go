package day02

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	ROCK    = 0
	PAPER   = 1
	SCISSOR = 2
)

const (
	ME   = 10
	ELF  = 11
	DRAW = 12
)

// Worst solution ever?
func getRockPaperScissor(s string) int {
	if s == "A" || s == "X" {
		return ROCK
	} else if s == "B" || s == "Y" {
		return PAPER
	} else {
		return SCISSOR
	}
}

func winner(f, s string) int {
	elf := getRockPaperScissor(f)
	me := getRockPaperScissor(s)
	if elf == ROCK {
		if me == PAPER {
			return ME
		}
		if me == SCISSOR {
			return ELF
		}
	} else if elf == PAPER {
		if me == ROCK {
			return ELF
		}
		if me == SCISSOR {
			return ME
		}
	} else if elf == SCISSOR {
		if me == ROCK {
			return ME
		}
		if me == PAPER {
			return ELF
		}
	}
	return DRAW
}

func calculateOneGame(elf, me string) (point int) {
	mine := getRockPaperScissor(me)
	if mine == ROCK {
		point += 1
	} else if mine == PAPER {
		point += 2
	} else {
		point += 3
	}

	who := winner(elf, me)
	if who == ME {
		point += 6
	} else if who == DRAW {
		point += 3
	}

	return
}

func scanGame(fileName string) (sum int, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, fmt.Errorf("unable to open file `%s`", fileName)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		game := strings.Split(scanner.Text(), " ")
		sum += calculateOneGame(game[0], game[1])
	}

	return
}
