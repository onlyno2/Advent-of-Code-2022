package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func exit(message string, err error) {
	println(message)
	panic(err)
}

func getAssignmentRange(inputRange string) (int, int) {
	_inputRange := strings.Split(inputRange, "-")

	begin, err := strconv.Atoi(_inputRange[0])
	if err != nil {
		exit("Invalid begin", err)
	}

	end, err := strconv.Atoi(_inputRange[1])
	if err != nil {
		exit("Invalid end", err)
	}

	return begin, end
}

// Can use struct to reduce the number of arguments
func fullyOverlapRange(begin1, end1, begin2, end2 int) bool {
	if begin1 <= begin2 && end1 >= end2 { // range 1 contains range 2
		return true
	} else if begin2 <= begin1 && end2 >= end1 { // range 2 contains range 1
		return true
	} else {
		return false
	}
}

// Can use struct to reduce the number of arguments
func overlapRange(begin1, end1, begin2, end2 int) bool {
	if begin1 <= begin2 && end1 >= begin2 {
		return true
	}

	if begin1 >= begin2 && end2 >= begin1 {
		return true
	}

	return false
}

func part1() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		exit("No input file", err)
	}

	scanner := bufio.NewScanner(readFile)
	total := 0

	for scanner.Scan() {
		assignments := strings.Split(scanner.Text(), ",")
		firstAssignmentBegin, firstAssignmentEnd := getAssignmentRange(assignments[0])
		secondAssignmentBegin, secondAssignmentEnd := getAssignmentRange(assignments[1])

		if fullyOverlapRange(firstAssignmentBegin, firstAssignmentEnd, secondAssignmentBegin, secondAssignmentEnd) {
			total += 1
		}
	}

	println(total)
}

func part2() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		exit("No input file", err)
	}

	scanner := bufio.NewScanner(readFile)
	total := 0

	for scanner.Scan() {
		assignments := strings.Split(scanner.Text(), ",")
		firstAssignmentBegin, firstAssignmentEnd := getAssignmentRange(assignments[0])
		secondAssignmentBegin, secondAssignmentEnd := getAssignmentRange(assignments[1])

		if overlapRange(firstAssignmentBegin, firstAssignmentEnd, secondAssignmentBegin, secondAssignmentEnd) {
			total += 1
		}
	}

	println(total)
}
