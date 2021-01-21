package main

import (
	"fmt"
	"os"
	"strconv"
)

func sieve(max int) []bool {
	a := make([]bool, max+1)

	for i := 0; i <= max; i++ {
		a[i] = true
	}

	for p := 2; p*p <= max; p++ {
		if a[p] == true {
			for idx := p * p; idx <= max; idx += p {
				a[idx] = false
			}
		}
	}

	return a
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("[usage error] - please provide an upper limit.")
		fmt.Println("Try again, such as: ~ ./sieve 123")
		os.Exit(-1)
	}

	max, _ := strconv.Atoi(args[0])
	fmt.Printf("[Sieve of Eratosthenes] - Find All Primes To An Upper Bound Of %v.\n", max)

	res := sieve(max)

	fmt.Print("Found all prime numbers up to max: ")
	for i := 2; i < len(res); i++ {
		if res[i] == true {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Print("\n\n")
}
