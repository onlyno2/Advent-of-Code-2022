package main

import (
	"bufio"
	"fmt"
	"math"
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

func loadDirectorySize() map[string]int {
	readFile, err := os.Open("input.txt")
	if err != nil {
		exit("No input file", err)
	}

	path := []string{}
	dirSizeMap := make(map[string]int)

	scanner := bufio.NewScanner(readFile)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		if tokens[1] == "cd" {
			if tokens[2] == ".." {
				path = path[:(len(path) - 1)]
			} else {
				path = append(path, tokens[2])
			}
		} else if tokens[1] == "ls" {
			continue
		} else {
			if tokens[0] == "dir" {
				continue
			}
			size, err := strconv.Atoi(tokens[0])
			if err != nil {
				exit("Invalid file size", err)
			}

			for i := 0; i < len(path)+1; i++ {
				dirSizeMap[strings.Join(path[:i], "/")] += size
			}
		}
	}

	return dirSizeMap
}

func part1() {
	total := 0
	dirSizeMap := loadDirectorySize()

	for _, dirSize := range dirSizeMap {
		if dirSize <= 100000 {
			total += dirSize
		}
	}

	fmt.Println(total)
}

func part2() {
	smallestSpaceToFree := math.Inf(1)
	dirSizeMap := loadDirectorySize()
	currentUsed := dirSizeMap["/"]
	maxUsed := 70000000 - 30000000
	needToFree := currentUsed - maxUsed

	for _, dirSize := range dirSizeMap {
		if dirSize >= needToFree {
			smallestSpaceToFree = math.Min(smallestSpaceToFree, float64(dirSize))
		}
	}

	fmt.Println(smallestSpaceToFree)
}
