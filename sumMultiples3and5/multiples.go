/* template Go */
package main

import (
	"fmt"
)

func main() {
	var sum int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
			fmt.Println(i, "is a multiple of 3 or 5.", sum, "is the sum of all multiples so far.")
		}
	}
}
