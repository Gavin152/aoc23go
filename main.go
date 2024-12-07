package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Gavin152/aoc23go/internal/filereader"
)

func main() {

	filePath := "example"
	// filePath := "data"

	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	possible := 0

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		tmp := strings.Split(line, ": ")
		game := tmp[0]
		rounds := strings.Split(tmp[1], "; ")
		gamePossible := true
		for _, round := range rounds {
			trimmed := strings.TrimSpace(round)
			colours := strings.Split(trimmed, ", ")
			for _, colour := range colours {
				split := strings.Split(colour, " ")
				count, _ := strconv.Atoi(split[0])
				color := split[1]
				isPossible := true
				if color == "red" && count > maxRed {
					isPossible = false
				}
				if color == "green" && count > maxGreen {
					isPossible = false
				}
				if color == "blue" && count > maxBlue {
					isPossible = false
				}
				if !isPossible {
					gamePossible = false
					break
				}
			}
		}
		if gamePossible {
			gs := strings.Split(game, " ")
			gameId, _ := strconv.Atoi(gs[1])
			possible += gameId
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Possible games: %d\n", possible)
}
