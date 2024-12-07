package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Gavin152/aoc23go/internal/filereader"
)

func main() {

	// filePath := "example"
	filePath := "data"

	sum := 0

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		minRed := 0
		minGreen := 0
		minBlue := 0
		tmp := strings.Split(line, ": ")
		rounds := strings.Split(tmp[1], "; ")
		for _, round := range rounds {
			trimmed := strings.TrimSpace(round)
			colours := strings.Split(trimmed, ", ")
			for _, colour := range colours {
				split := strings.Split(colour, " ")
				count, _ := strconv.Atoi(split[0])
				color := split[1]
				if color == "red" && count > minRed {
					minRed = count
				}
				if color == "green" && count > minGreen {
					minGreen = count
				}
				if color == "blue" && count > minBlue {
					minBlue = count
				}
			}
		}
		sum += minRed * minGreen * minBlue
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Sum of possible games: %d\n", sum)
}
