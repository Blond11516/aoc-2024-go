package main

import (
	"aoc2024/helpers"
	day1 "aoc2024/solvers/1"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	day := GetDay()
	part := GetPart()

	input := helpers.ReadFile("inputs/1.txt")

	result := solvers[day][part-1](input)

	fmt.Println(result)
}

type SolverFunc = func(string) int

var solvers = map[int]([2]SolverFunc){
	1: [2]SolverFunc{day1.Part1, day1.Part2},
}

func GetDay() int {
	day, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	return day
}

func GetPart() int {
	part, err := strconv.Atoi(os.Args[2])

	if err != nil {
		return 1
	}

	return part
}
