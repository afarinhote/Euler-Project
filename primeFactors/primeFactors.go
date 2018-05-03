package main

import (
	"fmt"
	"time"
)

func main() {
	var numberToFactor int
	var primeFactors []int
	factor := 2
	fmt.Print("Enter number to check prime factorization: ")
	fmt.Scanln(&numberToFactor)
	for numberToFactor/factor != 1 {
		if numberToFactor%factor == 0 {
			numberToFactor = numberToFactor / factor
			primeFactors = append(primeFactors, factor)
		} else {
			factor++
		}
	}
	primeFactors = append(primeFactors, numberToFactor)
	fmt.Println(primeFactors)
	time.Sleep(10 * time.Second)
}

/*		old code(very inefficient, kinda works but not really)
func main() {
	var prime, primeMultiplied int
	primeMultiplied = 1
	var allPrimeFactors, primeFactors []int
	fmt.Print("Enter number to check prime factorization: ")
	fmt.Scanln(&prime)
	for index := 0; index < prime; index++ {
		if checkIfPrime(index) {
			allPrimeFactors = append(allPrimeFactors, index)
		}
	}
	for primeMultiplied < prime {
		fmt.Println("why1")
		for k := 0; k < len(allPrimeFactors); k++ {
			if prime%allPrimeFactors[k] == 0 {
				primeFactors = append(primeFactors,allPrimeFactors[k])
				if k != 0 {prime = prime / allPrimeFactors[k]}
				primeMultiplied = primeMultiplied*allPrimeFactors[k]
				fmt.Println(allPrimeFactors[k])
			}
		}
		fmt.Println(primeFactors)
	}
}
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
}*/
