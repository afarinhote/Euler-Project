package main

import (
	"fmt"
)

func main() {
	var sum int
	var i int
	for i < 2000000 {
		if checkIfPrime(i) {
			sum += i
		}
		i++
	}
	fmt.Println(sum)
}

/*Repurposed code from failed Prime factorization attempt*/
func checkIfPrime(x int) bool {
	var checker int
	var check bool
	for i := 2; i < x; i++ {
		if x%i != 0 {
			checker = checker + 1
		}
	}
	if checker+2 == x {
		check = true
	} else {
		check = false
	}
	return check

}
