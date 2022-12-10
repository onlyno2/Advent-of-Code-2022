package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
	fmt.Println()
}

func part1() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		panic("Invalid input file")
	}

	scanner := bufio.NewScanner(readFile)
	x := 1
	currentCycle := 1
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")

		if currentCycle%40 == 20 {
			ans += currentCycle * x
		}
		if command[0] == "noop" {
			currentCycle += 1
		} else if command[0] == "addx" {
			value, _ := strconv.Atoi(command[1])

			currentCycle += 1
			if currentCycle%40 == 20 {
				ans += currentCycle * x
			}
			currentCycle += 1
			x += value
		}
	}

	fmt.Println(ans)
}

func draw(register, cycle int) {
	rowPixel := cycle%40 + 1
	if rowPixel >= register && rowPixel <= register+2 {
		print("#")
	} else {
		print(" ")
	}
}

func part2() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		panic("Invalid input file")
	}

	crtWidth := 40
	x := 1
	currentCycle := 0
	scanner := bufio.NewScanner(readFile)
	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")
		if currentCycle%crtWidth == 0 {
			fmt.Println()
		}

		draw(x, currentCycle)
		if command[0] == "noop" {
			currentCycle += 1
		} else if command[0] == "addx" {
			value, _ := strconv.Atoi(command[1])

			currentCycle += 1
			if currentCycle%crtWidth == 0 {
				fmt.Println()
			}
			draw(x, currentCycle)
			currentCycle += 1
			x += value
		}
	}
}
