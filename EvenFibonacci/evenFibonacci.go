package main

import (
	"fmt"
)

func main() {
	var sum, oldNumber, holder, evenTotal int
	sum, oldNumber = 1, 0
	for sum < 4000000 {
		holder = sum
		sum = oldNumber + sum
		oldNumber = holder
		if sum%2 == 0 {
			evenTotal = sum + evenTotal
			fmt.Println(evenTotal)
		}
	}
}
