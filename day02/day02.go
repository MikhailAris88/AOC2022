package day02

import (
	"AOC2022/src"
	"fmt"
)

func DayTwo() error {
	input, err := src.GetData()
	if err != nil {
		return err
	}
	data, err := input.Strings(2022, 2)
	if err != nil {
		return err
	}
	partOne(data)
	patTwo(data)

	return nil
}
func partOne(data []string) {
	var points int
	var sum int
	for _, v := range data {
		if string(v[0]) == "A" {
			switch {
			case string(v[2]) == "X":
				points = 1 + 3
			case string(v[2]) == "Y":
				points = 2 + 6
			case string(v[2]) == "Z":
				points = 3 + 0
			}
		}
		if string(v[0]) == "B" {
			switch {
			case string(v[2]) == "X":
				points = 1 + 0
			case string(v[2]) == "Y":
				points = 2 + 3
			case string(v[2]) == "Z":
				points = 3 + 6
			}
		}
		if string(v[0]) == "C" {
			switch {
			case string(v[2]) == "X":
				points = 1 + 6
			case string(v[2]) == "Y":
				points = 2 + 0
			case string(v[2]) == "Z":
				points = 3 + 3
			}
		}
		sum += points
	}
	fmt.Println("part one", sum)
}
func patTwo(data []string) {
	var points int
	var sum int
	for _, v := range data {
		if string(v[0]) == "A" {
			switch {
			case string(v[2]) == "X":
				points = 3 + 0
			case string(v[2]) == "Y":
				points = 1 + 3
			case string(v[2]) == "Z":
				points = 2 + 6
			}
		}
		if string(v[0]) == "B" {
			switch {
			case string(v[2]) == "X":
				points = 1 + 0
			case string(v[2]) == "Y":
				points = 2 + 3
			case string(v[2]) == "Z":
				points = 3 + 6
			}
		}
		if string(v[0]) == "C" {
			switch {
			case string(v[2]) == "X":
				points = 2 + 0
			case string(v[2]) == "Y":
				points = 3 + 3
			case string(v[2]) == "Z":
				points = 1 + 6
			}
		}
		sum += points
	}
	fmt.Println("part two", sum)
}
