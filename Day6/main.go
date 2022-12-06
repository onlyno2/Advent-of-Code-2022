package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	part1()
	part2()
}

func hasDuplicateCharacters(chars []rune) bool {
	charMap := make(map[rune]bool)

	for _, c := range chars {
		if charMap[c] {
			return true
		}

		charMap[c] = true
	}

	return false
}

func findMarker(length int) {
	readFile, err := os.Open("input.txt")
	if err != nil {
		exit("No input file", err)
	}

	reader := bufio.NewReader(readFile)
	chars := []rune{}

	for i := 1; i <= length; i++ {
		char, _, _ := reader.ReadRune()
		chars = append(chars, char)
	}

	if !hasDuplicateCharacters(chars) {
		println(length)
		return
	}

	counter := length + 1
	for {
		char, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}

		chars = append(chars, char)[1:]

		if !hasDuplicateCharacters(chars) {
			break
		}

		counter++
	}

	println(counter)
}

func part1() {
	findMarker(4)
}

func part2() {
	findMarker(14)
}

func exit(message string, err error) {
	println(message)
	panic(err)
}
