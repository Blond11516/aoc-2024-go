package day3

import (
	"aoc2024/helpers"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) int {
	regex := helpers.Safe(regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`))
	matches := regex.FindAll([]byte(input), -1)

	sum := 0
	for _, match := range matches {
		stringMatch := string(match)
		withoutSuffix, _ := strings.CutPrefix(stringMatch, "mul(")
		trimmed, _ := strings.CutSuffix(withoutSuffix, ")")

		rawNumbers := strings.Split(trimmed, ",")
		firstNumber := helpers.Safe(strconv.Atoi(rawNumbers[0]))
		secondNumber := helpers.Safe(strconv.Atoi(rawNumbers[1]))

		sum += firstNumber * secondNumber
	}

	return sum
}

func Part2(input string) int {
	regex := helpers.Safe(regexp.Compile(`mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`))
	matches := regex.FindAll([]byte(input), -1)

	sum := 0
	enabled := true
	for _, match := range matches {
		switch string(match) {
		case "don't()":
			enabled = false
		case "do()":
			enabled = true
		default:
			if enabled {
				stringMatch := string(match)
				withoutSuffix, _ := strings.CutPrefix(stringMatch, "mul(")
				trimmed, _ := strings.CutSuffix(withoutSuffix, ")")

				rawNumbers := strings.Split(trimmed, ",")
				firstNumber := helpers.Safe(strconv.Atoi(rawNumbers[0]))
				secondNumber := helpers.Safe(strconv.Atoi(rawNumbers[1]))

				sum += firstNumber * secondNumber
			}
		}
	}

	return sum
}
