package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1()
	part2()
}

func exit(message string, err error) {
	println(message)
	panic(err)
}

func loadTreeMap() ([][]int, int, int) {
	readFile, err := os.Open("input.txt")
	if err != nil {
		exit("No input file", err)
	}

	treeMap := [][]int{}

	scanner := bufio.NewScanner(readFile)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		treeMap = append(treeMap, []int{})

		for _, heightRune := range line {
			treeMap[row] = append(treeMap[row], int(heightRune)-'0')
		}

		row++
	}
	col := len(treeMap[0])

	return treeMap, row, col
}

func isHigest(currentHeight int, heights []int) bool {
	for _, height := range heights {
		if currentHeight <= height {
			return false
		}
	}

	return true
}

func getCol(treeMap [][]int, startRow, endRow, col int) []int {
	result := []int{}

	for i := startRow; i < endRow; i++ {
		result = append(result, treeMap[i][col])
	}

	return result
}

func part1() {
	treeMap, row, col := loadTreeMap()

	ans := 0
	for i := 0; i < row; i++ {
		if i == 0 || i == row-1 {
			ans += col
			continue
		}

		for j := 0; j < col; j++ {
			if j == 0 || j == col-1 {
				ans += 1
				continue
			}

			currentHeight := treeMap[i][j]
			isVisible := isHigest(currentHeight, treeMap[i][j+1:]) ||
				isHigest(currentHeight, treeMap[i][:j]) ||
				isHigest(currentHeight, getCol(treeMap, i+1, row, j)) ||
				isHigest(currentHeight, getCol(treeMap, 0, i, j))

			if isVisible {
				ans += 1
			}
		}
	}

	fmt.Println(ans)
}

func calScore(treeMap [][]int, i, j, row, col int) int {
	currentHeight := treeMap[i][j]

	leftScore := 0
	rightScore := 0
	topScore := 0
	botScore := 0

	// Calculate left score
	for tmp := j - 1; tmp >= 0; tmp-- {
		leftScore++

		if currentHeight <= treeMap[i][tmp] {
			break
		}
	}

	// Calculate right score
	for tmp := j + 1; tmp < col; tmp++ {
		rightScore++

		if currentHeight <= treeMap[i][tmp] {
			break
		}
	}

	// Calculate top score
	top := getCol(treeMap, 0, i, j)
	for tmp := len(top) - 1; tmp >= 0; tmp-- {
		if top[tmp] < currentHeight {
			topScore++
		} else {
			topScore++
			break
		}
	}

	// Calculate bottom score
	bot := getCol(treeMap, i+1, row, j)
	for tmp := 0; tmp < len(bot); tmp++ {
		if bot[tmp] < currentHeight {
			botScore++
		} else {
			botScore++
			break
		}
	}

	return leftScore * rightScore * topScore * botScore
}

func part2() {
	treeMap, row, col := loadTreeMap()

	ans := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			score := calScore(treeMap, i, j, row, col)
			if ans < score {
				ans = score
			}
		}
	}

	fmt.Println(ans)
}
