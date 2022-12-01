package main

import (
	"AOC2022/day01"
	"fmt"
)

func main() {
	answer, err := day01.DayOne()
	if err != nil {
		fmt.Println("Error data")
	}
	fmt.Println(answer)
}
