package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, i, j float64
	for i = 1; i < 999; i++ {
		for j = 1; j < 998; j++ {
			a = math.Sqrt(i*i + j*j)
			b = a + i + j
			if b == 1000 {
				b = a * i * j
				fmt.Println("Product:", b, "a:", i, "b:", j, "c:", a)
				break
			}
		}
	}
}
