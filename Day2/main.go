package main

import (
	"bufio"
	"os"
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

var movePoints = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var outcomePoints = map[string]int{
	"win":  6,
	"draw": 3,
	"lost": 0,
}

func getOutcomeScore(outcome string) int {
	return outcomePoints[outcome]
}

func getAppropriateMove(opponentMove string, outcome string) string {
	switch opponentMove {
	case "rock":
		if outcome == "win" {
			return "paper"
		} else if outcome == "draw" {
			return "rock"
		} else if outcome == "lose" {
			return "scissors"
		} else {
			panic("Invalid outcome")
		}
	case "paper":
		if outcome == "win" {
			return "scissors"
		} else if outcome == "draw" {
			return "paper"
		} else if outcome == "lose" {
			return "rock"
		} else {
			panic("Invalid outcome")
		}
	case "scissors":
		if outcome == "win" {
			return "rock"
		} else if outcome == "draw" {
			return "scissors"
		} else if outcome == "lose" {
			return "paper"
		} else {
			panic("Invalid outcome")
		}
	default:
		panic("Invalid move")
	}
}

func compareMove(move1 string, move2 string) string {
	switch move1 {
	case "rock":
		if move2 == "rock" {
			return "draw"
		} else if move2 == "paper" {
			return "win"
		} else if move2 == "scissors" {
			return "lose"
		} else {
			panic("Invalid move")
		}
	case "paper":
		if move2 == "rock" {
			return "lose"
		} else if move2 == "paper" {
			return "draw"
		} else if move2 == "scissors" {
			return "win"
		} else {
			panic("Invalid move")
		}
	case "scissors":
		if move2 == "rock" {
			return "win"
		} else if move2 == "paper" {
			return "lose"
		} else if move2 == "scissors" {
			return "draw"
		} else {
			panic("Invalid move")
		}
	default:
		panic("Invalid move")
	}
}

func part1() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		exit("No input file", err)
	}

	availableMoves := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	getScore := func(opponentMove string, myMove string) int {
		outcome := compareMove(opponentMove, myMove)

		return getOutcomeScore(outcome)
	}

	scanner := bufio.NewScanner(readFile)
	score := 0
	for scanner.Scan() {
		round := scanner.Text()
		moves := strings.Fields(round)
		opponentMove := availableMoves[moves[0]]
		myMove := availableMoves[moves[1]]

		score += getScore(opponentMove, myMove) + movePoints[myMove]
	}

	println(score)
}

func part2() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		exit("No input file", err)
	}

	availableMoves := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}
	availableOutcomes := map[string]string{
		"X": "lose",
		"Y": "draw",
		"Z": "win",
	}

	scanner := bufio.NewScanner(readFile)
	totalScore := 0
	for scanner.Scan() {
		round := scanner.Text()
		roundResult := strings.Fields(round)
		opponentMove := availableMoves[roundResult[0]]
		outcome := availableOutcomes[roundResult[1]]

		totalScore += getOutcomeScore(outcome) + movePoints[getAppropriateMove(opponentMove, outcome)]
	}

	println(totalScore)
}
