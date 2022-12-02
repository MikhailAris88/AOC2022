package day01

import (
	"AOC2022/src"
	"fmt"
	"sort"
	"strconv"
)

func DayOne() (string, error) {
	input, err := src.GetData()
	if err != nil {
		return "", err
	}
	data, err := input.Strings(2022, 1)
	if err != nil {
		return "", err
	}
	var allElfCalories []int
	var indexELf int
	var oneElfCalories int
	for _, val := range data {
		colories, err := strconv.Atoi(val)
		if err != nil {
			allElfCalories = append(allElfCalories, oneElfCalories)
			indexELf++
			oneElfCalories = 0
		}
		oneElfCalories += colories
	}
	sort.Slice(allElfCalories, func(i, j int) bool {
		return allElfCalories[i] > allElfCalories[j]
	})
	var bestThree int
	for i := 0; i < 3; i++ {
		bestThree += allElfCalories[i]
	}
	answer := fmt.Sprintf("max one elf calories = %d, 3 best elf = %d", allElfCalories[0], bestThree)
	return answer, nil
}
