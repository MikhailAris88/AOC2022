package day03

import (
	"AOC2022/src"
	"fmt"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func DayThree() error {
	input, err := src.GetData()
	if err != nil {
		return err
	}
	data, err := input.Strings(2022, 3)
	if err != nil {
		return err
	}
	pOne := partOne(data)
	pTwo := partTwo(data)
	fmt.Println("Part one:", pOne, "Part two:", pTwo)
	return nil
}

func partOne(data []string) int {
	var sum int
	for _, str := range data {
		lenStr := len(str)
		fSub := str[:lenStr/2]
		sSub := str[len(fSub):lenStr]
		for _, s := range fSub {
			if strings.Contains(sSub, string(s)) {
				index := strings.IndexAny(alphabet, string(s)) + 1
				sum += index
				break
			}
		}
	}
	return sum
}

func partTwo(data []string) int {
	var sum int
	for i := 0; i < len(data); i += 3 {
		for _, v := range data[i] {
			if strings.Contains(data[i+1], string(v)) {
				if strings.Contains(data[i+2], string(v)) {
					sum += 1 + strings.Index(alphabet, string(v))
					break
				}
			}
		}
	}
	return sum
}
