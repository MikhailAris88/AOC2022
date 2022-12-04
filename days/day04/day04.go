package day04

import (
	"AOC2022/src"
	"fmt"
	"strconv"
	"strings"
)

func DayFour() error {
	input, err := src.GetData()
	if err != nil {
		return err
	}
	data, err := input.Strings(2022, 4)
	if err != nil {
		return err
	}
	pOne := partOne(data)
	pTwo := partTwo(data)
	fmt.Println("part one answer:", pOne, "part two answer:", pTwo)
	return nil
}

func partOne(data []string) int {
	match := 0
	for i := 0; i < len(data); i++ {
		line := strings.Split(data[i], ",")
		firstPair := strings.Split(line[0], "-")
		a1, _ := strconv.Atoi(firstPair[0])
		b1, _ := strconv.Atoi(firstPair[1])
		secondPair := strings.Split(line[1], "-")
		a2, _ := strconv.Atoi(secondPair[0])
		b2, _ := strconv.Atoi(secondPair[1])
		if a1 <= a2 && b1 >= b2 || a1 >= a2 && b1 <= b2 {
			match++
		}
	}
	return match
}

func partTwo(data []string) int {
	match := 0
	for i := 0; i < len(data); i++ {
		line := strings.Split(data[i], ",")
		firstPair := strings.Split(line[0], "-")
		a1, _ := strconv.Atoi(firstPair[0])
		b1, _ := strconv.Atoi(firstPair[1])
		secondPair := strings.Split(line[1], "-")
		a2, _ := strconv.Atoi(secondPair[0])
		b2, _ := strconv.Atoi(secondPair[1])
		if a1 <= b2 && b1 >= a2 {
			match++
		}
	}
	return match
}
