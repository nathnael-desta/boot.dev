package main

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	for n := 2 ; n < max + 1; n++ {
		if n == 2 {
			continue
		}
		if n % 2 == 0 {
			continue
		} 
		prime := true
		for i := 3.0; i <= math.Sqrt(float64(n)); i += 2 {
			if n % int(i) == 0 {
				prime = false
				break 
			}
		} 
		if prime {
			fmt.Println(n)
		}


	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
