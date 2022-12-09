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
}

func exit(message string, err error) {
	println(message)
	panic(err)
}

type position struct {
	x int
	y int
}

var xTransition = map[string]int{
	"L":  -1, // move left
	"R":  1,  // move right
	"U":  0,  // move up
	"D":  0,  // move down,
	"LU": -1, // move left up
	"RU": 1,  // move right up
	"LD": -1, // move left down
	"RD": 1,  // move right down
}

var yTransition = map[string]int{
	"L":  0,  // move left
	"R":  0,  // move right
	"U":  -1, // move u
	"D":  1,  // move down
	"LU": -1, // move left up
	"RU": -1, // move right up
	"LD": 1,  // move left down
	"RD": 1,  // move right down
}

func applyMove(move []string, head *position, tail *position, visited map[string]bool) {
	steps, _ := strconv.Atoi(move[1])

	for i := 0; i < steps; i++ {
		head.x += xTransition[move[0]]
		head.y += yTransition[move[0]]
		direction := tailMoveDirection(head, tail)
		tail.x += xTransition[direction]
		tail.y += yTransition[direction]

		visitKey := fmt.Sprintf("%d:%d", tail.x, tail.y)
		visited[visitKey] = true
	}
}

func applyMove2(move []string, rope []*position, visited map[string]bool) {
	steps, _ := strconv.Atoi(move[1])

	for i := 0; i < steps; i++ {
		rope[0].x += xTransition[move[0]]
		rope[0].y += yTransition[move[0]]

		for j := 1; j < 10; j++ {
			direction := tailMoveDirection(rope[j-1], rope[j])

			rope[j].x += xTransition[direction]
			rope[j].y += yTransition[direction]

			if j == 9 {
				visitKey := fmt.Sprintf("%d:%d", rope[j].x, rope[j].y)
				visited[visitKey] = true
			}
		}
	}
}

func tailMoveDirection(head, tail *position) string {
	if head.x == tail.x+2 && head.y == tail.y {
		return "R"
	}

	if head.x == tail.x-2 && head.y == tail.y {
		return "L"
	}

	if head.y == tail.y+2 && head.x == tail.x {
		return "D"
	}

	if head.y == tail.y-2 && head.x == tail.x {
		return "U"
	}

	if head.x == tail.x+2 && head.y == tail.y+1 ||
		head.x == tail.x+2 && head.y == tail.y+2 ||
		head.x == tail.x+1 && head.y == tail.y+2 {
		return "RD"
	}

	if head.x == tail.x+2 && head.y == tail.y-1 ||
		head.x == tail.x+2 && head.y == tail.y-2 ||
		head.x == tail.x+1 && head.y == tail.y-2 {
		return "RU"
	}

	if head.x == tail.x-2 && head.y == tail.y+1 ||
		head.x == tail.x-2 && head.y == tail.y+2 ||
		head.x == tail.x-1 && head.y == tail.y+2 {
		return "LD"
	}

	if head.x == tail.x-2 && head.y == tail.y-1 ||
		head.x == tail.x-2 && head.y == tail.y-2 ||
		head.x == tail.x-1 && head.y == tail.y-2 {
		return "LU"
	}

	return "None"
}

func part1() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		exit("No input file", err)
	}

	scanner := bufio.NewScanner(readFile)
	head := position{
		x: 0,
		y: 0,
	}
	tail := position{
		x: 0,
		y: 0,
	}

	visited := make(map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()
		move := strings.Split(line, " ")
		applyMove(move, &head, &tail, visited)
	}

	fmt.Println(len(visited))
}

func part2() {
	var rope []*position
	for i := 0; i < 10; i++ {
		rope = append(rope, &position{
			x: 0,
			y: 0,
		})
	}
	visited := map[string]bool{"0:0": true}

	readFile, err := os.Open("input.txt")
	if err != nil {
		exit("No input file", err)
	}

	scanner := bufio.NewScanner(readFile)

	for scanner.Scan() {
		line := scanner.Text()
		move := strings.Split(line, " ")
		applyMove2(move, rope, visited)
	}

	fmt.Println(len(visited))
}
