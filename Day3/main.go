package main

import (
	"bufio"
	"os"
)

func main() {
	part1()
	part2()
}

func getRucksacks(rucksacks string) (string, string) {
	return rucksacks[:(len(rucksacks) / 2)], rucksacks[(len(rucksacks) / 2):]
}

func getPiority(item rune) uint {
	if item >= 'a' && item <= 'z' {
		return uint(item - 96)
	} else if item >= 'A' && item <= 'Z' {
		return uint(item - 38)
	} else {
		panic("Invalid item")
	}
}

func part1() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		exit("No input file", err)
	}

	scanner := bufio.NewScanner(readFile)
	var total uint = 0

	for scanner.Scan() {
		rucksacks := scanner.Text()
		firstRucksack, secondRucksack := getRucksacks(rucksacks)
		itemMap := make(map[string]bool)

		for _, item := range firstRucksack {
			itemMap[string(item)] = true
		}

		for _, item := range secondRucksack {
			if itemMap[string(item)] {
				total += getPiority(item)
				break
			}
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
	var total uint = 0

	for scanner.Scan() {
		firstRucksack := scanner.Text()
		scanner.Scan()
		secondRucksack := scanner.Text()
		scanner.Scan()
		thirdRucksack := scanner.Text()

		itemMap := make(map[rune]int)
		for _, item := range firstRucksack {
			itemMap[item] = 1
		}

		for _, item := range secondRucksack {
			if itemMap[item] == 1 {
				itemMap[item] = 2
			}
		}

		for _, item := range thirdRucksack {
			if itemMap[item] == 2 {
				total += getPiority(item)
				break
			}
		}
	}

	println(total)
}

func exit(message string, err error) {
	println(message)
	panic(err)
}
