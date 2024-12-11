package day2

import (
	"aoc2024/helpers"
	"slices"
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

		if safe {
			safeCount++
		}
	}

	return safeCount
}

func Part2(input string) int {
	lines := helpers.Lines(input)

	safeCount := 0
	for _, line := range lines {
		rawLevels := strings.Split(line, " ")

		var levels []int
		for _, char := range rawLevels {
			levels = append(levels, helpers.Safe(strconv.Atoi(char)))
		}

		if isSafe(levels, false) {
			safeCount++
		}
	}

	return safeCount
}

func isSafe(levels []int, dampened bool) bool {
	var increasing bool
	increasingSet := false
	for i := range len(levels) - 1 {
		firstLevel := levels[i]
		secondLevel := levels[i+1]

		diff := firstLevel - secondLevel
		if diff < 0 {
			diff *= -1
		}
		if diff < 1 || diff > 3 {
			return isRecoverable(levels, i, dampened)
		}
		if !increasingSet {
			increasingSet = true
			increasing = firstLevel < secondLevel
		}

		if increasing && firstLevel > secondLevel {
			return isRecoverable(levels, i, dampened)
		}

		if !increasing && firstLevel < secondLevel {
			return isRecoverable(levels, i, dampened)
		}
	}

	return true
}

func isRecoverable(levels []int, i int, dampened bool) bool {
	if dampened {
		return false
	} else {
		return testDampenedLevels(levels, i)
	}
}

func testDampenedLevels(levels []int, i int) bool {
	previousRemoved := []int{}
	if i > 0 {
		previousRemoved = slices.Delete(slices.Clone(levels), i-1, i)
	}
	firstRemoved := slices.Delete(slices.Clone(levels), i, i+1)
	secondRemoved := slices.Delete(slices.Clone(levels), i+1, i+2)
	return (len(previousRemoved) > 0 && isSafe(previousRemoved, true)) || isSafe(firstRemoved, true) || isSafe(secondRemoved, true)
}
