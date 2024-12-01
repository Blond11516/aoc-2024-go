package day1

import (
	"aoc2024/helpers"
	"log"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) int {
	lines := helpers.Lines(input)
	leftNumbers, rightNumbers := ParseLists(lines)

	sum := 0
	for i := range len(leftNumbers) {
		diff := leftNumbers[i] - rightNumbers[i]
		if diff < 0 {
			diff *= -1
		}
		sum += diff
	}

	return sum
}

func Part2(input string) int {
	lines := helpers.Lines(input)
	leftNumbers, rightNumbers := ParseLists(lines)

	rightNumberFrequencies := map[int]int{}
	for _, number := range rightNumbers {
		freq, ok := rightNumberFrequencies[number]

		if ok {
			rightNumberFrequencies[number] = freq + 1
		} else {
			rightNumberFrequencies[number] = 1
		}
	}

	sum := 0
	for _, left := range leftNumbers {
		if freq, ok := rightNumberFrequencies[left]; ok {
			sum += left * freq
		}
	}

	return sum
}

func ParseLists(lines []string) ([]int, []int) {
	var leftNumbers []int
	var rightNumbers []int

	for _, line := range lines {
		chars := strings.Split(line, "   ")

		number, err := strconv.Atoi(chars[0])
		if err != nil {
			log.Fatal(err)
		}
		leftNumbers = append(leftNumbers, number)

		number, err = strconv.Atoi(chars[1])
		if err != nil {
			log.Fatal(err)
		}
		rightNumbers = append(rightNumbers, number)
	}
	slices.Sort(leftNumbers)
	slices.Sort(rightNumbers)
	return leftNumbers, rightNumbers
}
