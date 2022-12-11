package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	part1()
	part2()
}

type Monkey struct {
	startingItems []int
	operation     func(int) int
	testFunc      func(int) int
	testValue     int
}

func loadMonkeys() []Monkey {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	monkeys := make([]Monkey, len(split))
	for _, data := range split {
		var items, operation string
		var id, value, test, trueThrow, falseThrow int
		fmt.Sscanf(
			strings.NewReplacer(", ", ",", "* old", "^ 2").Replace(data),
			`Monkey %d:
			Starting items: %s
			Operation: new = old %s %d
			Test: divisible by %d
				If true: throw to monkey %d
				If false: throw to monkey %d`,
			&id, &items, &operation, &value, &test, &trueThrow, &falseThrow,
		)

		json.Unmarshal([]byte("["+items+"]"), &monkeys[id].startingItems)
		monkeys[id].operation = map[string]func(int) int{
			"+": func(oldValue int) int { return oldValue + value },
			"*": func(oldValue int) int { return oldValue * value },
			"^": func(oldValue int) int { return oldValue * oldValue },
		}[operation]
		monkeys[id].testFunc = func(v int) int {
			if v%test == 0 {
				return trueThrow
			}

			return falseThrow
		}
		monkeys[id].testValue = test
	}

	return monkeys
}

func inspect(monkeys []Monkey, rounds int, op func(int) int) {
	inspected := make([]int, len(monkeys))

	for i := 0; i < rounds; i++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.startingItems {
				item = op(monkey.operation(item))
				monkeys[monkey.testFunc(item)].startingItems = append(monkeys[monkey.testFunc(item)].startingItems, item)
				inspected[i]++
			}
			monkeys[i].startingItems = nil
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))
	fmt.Println(inspected[0] * inspected[1])
}

func part1() {
	inspect(loadMonkeys(), 20, func(i int) int { return i / 3 })
}

func part2() {
	monkeys := loadMonkeys()
	wl := 1
	for _, monkey := range monkeys {
		wl *= monkey.testValue
	}

	inspect(loadMonkeys(), 10000, func(i int) int { return i % wl })
}
