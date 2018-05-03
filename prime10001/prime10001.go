package main

import (
	"fmt"
)

func main() {
	var allPrimeFactors []int
	i := 0
	for len(allPrimeFactors) <= 10000 {
		if checkIfPrime(i) {
			allPrimeFactors = append(allPrimeFactors, i)
		}
		i++
	}
	fmt.Println(allPrimeFactors[10000])
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
