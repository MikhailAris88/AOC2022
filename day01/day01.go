package day01

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	calories int
}

func DayOne() (string, error) {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	s := strings.Split(string(data), "\n")
	elfs := make([]*Elf, 0)
	var colories int
	i := 1
	s = append(s, "")
	for _, val := range s {
		col, err := strconv.Atoi(val)
		colories += col
		if err != nil {
			e := &Elf{
				calories: colories,
			}
			elfs = append(elfs, e)
			i++
			colories = 0
			continue
		}
	}
	max := 0
	maxElf := new(Elf)
	for _, val := range elfs {
		if val.calories > max {
			max = val.calories
			maxElf = val
		}
	}

	sliceInts := make([]int, 0)
	for _, v := range elfs {
		sliceInts = append(sliceInts, v.calories)
	}

	sort.Slice(sliceInts, func(i, j int) bool {
		return sliceInts[i] > sliceInts[j]
	})

	return fmt.Sprintf("max colories = %d, best3 = %d", maxElf.calories, sliceInts[0]+sliceInts[1]+sliceInts[2]), nil
}
