package main

import (
	"aoc2024/helpers"
	day1 "aoc2024/solvers/1"
	day2 "aoc2024/solvers/2"
	day3 "aoc2024/solvers/3"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	day := GetDay()
	part := GetPart()

	input := helpers.ReadFile(fmt.Sprintf("inputs/%d.txt", day))

	result := solvers[day][part-1](input)

	fmt.Println(result)
}

type SolverFunc = func(string) int

var solvers = map[int]([2]SolverFunc){
	1: [2]SolverFunc{day1.Part1, day1.Part2},
	2: [2]SolverFunc{day2.Part1, day2.Part2},
	3: [2]SolverFunc{day3.Part1, day3.Part2},
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
