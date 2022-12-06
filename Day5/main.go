package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	println()
	part2()
	println()
}

func exit(message string, err error) {
	println(message)
	panic(err)
}

func loadStacks(scanner *bufio.Scanner) map[int][]rune {
	stacks := make(map[int][]rune)

	for i := 0; i <= 9; i++ {
		stacks[i] = []rune{}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		if !strings.Contains(line, "[") {
			continue
		}

		for pos, char := range line {
			if pos%4 == 1 && string(char) != " " {
				stacks[int(pos/4)+1] = append(stacks[int(pos/4)+1], char)
			}
		}
	}

	return stacks
}

func extractCommand(command string) (int, int, int) {
	replacer := strings.NewReplacer(
		"move ", "",
		"from ", "",
		"to ", "",
	)

	extractCommand := replacer.Replace(command)
	commandArgs := strings.Split(extractCommand, " ")
	move, err := strconv.Atoi(commandArgs[0])
	if err != nil {
		exit("Invalid command", err)
	}
	from, err := strconv.Atoi(commandArgs[1])
	if err != nil {
		exit("Invalid command", err)
	}
	to, err := strconv.Atoi(commandArgs[2])
	if err != nil {
		exit("Invalid command", err)
	}

	return move, from, to
}

func part1() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		exit("No input file", err)
	}

	scanner := bufio.NewScanner(readFile)
	cratesMap := loadStacks(scanner)

	for scanner.Scan() {
		command := scanner.Text()

		move, from, to := extractCommand(command)
		for i := 0; i < move; i++ {
			var pop rune
			pop, cratesMap[from] = cratesMap[from][0], cratesMap[from][1:]
			cratesMap[to] = append([]rune{pop}, cratesMap[to]...)
		}
	}

	for i := 1; i < 10; i++ {
		print(string(cratesMap[i][0]))
	}
}

func part2() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		exit("No input file", err)
	}

	scanner := bufio.NewScanner(readFile)
	cratesMap := loadStacks(scanner)

	for scanner.Scan() {
		command := scanner.Text()
		move, from, to := extractCommand(command)

		var pop []rune
		for i := 0; i < move; i++ {
			pop = append(pop, cratesMap[from][0])
			cratesMap[from] = cratesMap[from][1:]
		}
		cratesMap[to] = append(pop, cratesMap[to]...)
	}

	for i := 1; i < 10; i++ {
		print(string(cratesMap[i][0]))
	}
}
