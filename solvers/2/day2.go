package day2

import (
	"aoc2024/helpers"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) int {
	lines := helpers.Lines(input)

	safeCount := 0
	for _, line := range lines {
		rawLevels := strings.Split(line, " ")

		var levels []int
		for _, char := range rawLevels {
			levels = append(levels, helpers.Safe(strconv.Atoi(char)))
		}

		increasing := levels[0] < levels[1]

		safe := true
		for i := range len(levels) - 1 {
			firstIndex := i
			secondIndex := firstIndex + 1
			if increasing && levels[firstIndex] > levels[secondIndex] {
				safe = false
				break
			}

			if !increasing && levels[firstIndex] < levels[secondIndex] {
				safe = false
				break
			}

			diff := levels[firstIndex] - levels[secondIndex]
			if diff < 0 {
				diff *= -1
			}
			if diff < 1 || diff > 3 {
				safe = false
				break
			}
		}

		if !safe {
			fmt.Println(levels, safe)
		}
		if safe {
			safeCount++
		}
	}

	return safeCount
}
