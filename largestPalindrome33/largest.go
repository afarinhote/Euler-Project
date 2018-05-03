package main

import (
	"fmt"
	"strconv"
)

func main() {
	number1, number2 := 1, 1
	var multiplied, biggest int
	for number1 < 1000 {
		for number2 = number1; number2 < 1000; number2++ {
			multiplied = number1 * number2
			if len(strconv.Itoa(multiplied)) == 6 {
				palinCheck := strconv.Itoa(multiplied)
				if palinCheck[0] == palinCheck[5] && palinCheck[1] == palinCheck[4] && palinCheck[2] == palinCheck[3] {
					if multiplied > biggest {
						biggest = multiplied
					}
				}
			}
		}
		number1++
	}
	fmt.Println(biggest)
}
