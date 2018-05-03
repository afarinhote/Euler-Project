package main

import (
	"fmt"
)

func main() {
	var sum, sumOfSquares int
	for i := 1; i < 101; i++ {
		sum += i
		sumOfSquares += i * i
	}
	fmt.Println(sum*sum - sumOfSquares)
}
