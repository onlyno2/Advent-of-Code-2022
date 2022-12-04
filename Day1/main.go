package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		exit("No input file", err)
	}

	scanner := bufio.NewScanner(readFile)
	highestCallories := 0
	secondHighestCallories := 0
	thirdHighestCallories := 0
	currentElfCallories := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 {
			if currentElfCallories < thirdHighestCallories {
				currentElfCallories = 0
				continue
			}
			thirdHighestCallories = currentElfCallories

			if currentElfCallories < secondHighestCallories {
				currentElfCallories = 0
				continue
			}
			thirdHighestCallories = secondHighestCallories
			secondHighestCallories = currentElfCallories

			if currentElfCallories < highestCallories {
				currentElfCallories = 0
				continue
			}
			secondHighestCallories = highestCallories
			highestCallories = currentElfCallories

			currentElfCallories = 0
		} else {
			callories, err := strconv.Atoi(line)
			if err != nil {
				exit("Invalid callories", err)
			}

			currentElfCallories += callories
		}
	}

	println(highestCallories)
	println(highestCallories + secondHighestCallories + thirdHighestCallories)
}

func exit(message string, err error) {
	println(message)
	panic(err)
}
