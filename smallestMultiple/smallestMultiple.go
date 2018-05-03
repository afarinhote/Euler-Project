package main

import (
	"fmt"
)

/*terribly Inefficient*/
func main() {
	numberPrint, count := 0, 0
	for j := 250000000; j > 0; j-- {
		count = 0
		for i := 1; i < 21; i++ {
			if j%i == 0 {
				count++
			}
			if count == 20 {
				fmt.Println(count, j)
				numberPrint = j
			}
		}
	}
	fmt.Println(numberPrint)
}
