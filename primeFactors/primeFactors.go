package main

import "fmt"

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
		for k := 1; k < len(allPrimeFactors)+1; k++ {
			if prime%allPrimeFactors[k-1] == 0 {
				primeFactors = append(primeFactors, allPrimeFactors[k-1])
				prime = prime / allPrimeFactors[k-1]
				primeMultiplied = primeMultiplied * allPrimeFactors[k-1]
				fmt.Println(allPrimeFactors[k-1])
				fmt.Println(primeMultiplied)
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
}
