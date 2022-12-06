package day06

import (
	"AOC2022/src"
	"fmt"
	"strings"
)

func DaySix() error {
	input, err := src.GetData()
	if err != nil {
		return err
	}
	data, err := input.Strings(2022, 6)
	if err != nil {
		return err
	}
	//для первой части счетчик стартуем с 4 шаг 3, для второй 14 и 13 соответственно
	pOne := puzzle(data, 4, 3)
	pTwo := puzzle(data, 14, 13)
	fmt.Println("part one answer:", pOne, "part two answer:", pTwo)
	return nil
}

func puzzle(data []string, count, step int) string {
	inputStr := strings.Join(data, "")
	for i := 1; i < len(inputStr)-3; i++ {
		match := false
		s := inputStr[i-1 : step+i]
		for _, v := range inputStr[i-1 : step+i] {
			f := string(v)
			e := strings.Replace(s, f, "", 1)
			if strings.Contains(e, f) {
				match = true
				break
			}
		}
		if match {
			count++
			continue
		} else {
			break
		}
	}
	return fmt.Sprintf("%d", count)

}
