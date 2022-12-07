package day05

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func DayFive() {
	input, _ := os.ReadFile("2022_5.txt")
	split := strings.Split(string(input), "\n\n")
	layout := strings.Split(split[0], "\n")
	stacks := make(map[int][]rune)
	stackKeys := []int{}
	for i := len(layout) - 1; i >= 0; i-- {
		if i == len(layout)-1 {
			for _, k := range strings.Split(strings.TrimSpace(layout[i]), "   ") {
				key, _ := strconv.Atoi(k)
				stacks[key] = []rune{}
				stackKeys = append(stackKeys, key)
			}
		} else {
			for i, c := range layout[i] {
				if !strings.ContainsAny(string(c), " []") {
					key := int(math.Ceil(float64(i) / 4))
					stacks[key] = append(stacks[key], c)
				}
			}
		}

	}
	moves := strings.Split(split[1], "\n")
	partOne(moves, stacks, stackKeys)
	partTwo(moves, stacks, stackKeys)
}

func partOne(moves []string, stacks map[int][]rune, stackKeys []int) string {
	for _, move := range moves {
		var amount int
		var from int
		var to int
		fmt.Sscanf(move, "move %d from %d to %d", &amount, &from, &to)
		source := stacks[from]
		destination := stacks[to]
		for i := 0; i < amount; i++ {
			n := len(source) - 1 // Top element
			destination = append(destination, source[n])
			source = source[:n] // Pop
		}
		stacks[from] = source
		stacks[to] = destination
	}
	pOne := ""
	for _, k := range stackKeys {
		pOne += string(stacks[k][len(stacks[k])-1])
	}
	return fmt.Sprintf("part one answer:%s", pOne)

}

func partTwo(moves []string, stacks map[int][]rune, stackKeys []int) string {
	for _, move := range moves {
		var amount int
		var from int
		var to int
		fmt.Sscanf(move, "move %d from %d to %d", &amount, &from, &to)
		source := stacks[from]
		destination := stacks[to]
		n := len(source) - amount
		for _, r := range source[n:] {
			destination = append(destination, r)
		}
		source = source[:n]
		stacks[from] = source
		stacks[to] = destination
	}
	pTwo := ""
	for _, k := range stackKeys {
		pTwo += string(stacks[k][len(stacks[k])-1])
	}
	return fmt.Sprintf("part two answer:%s", pTwo)
}
