//Problem 37
//Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

package projecteuler

import (
	"fmt"
	"primes"
	"strconv"
)

func isTruncable(prime int, primes map[int]int) bool {
	stra := strconv.Itoa(prime)
	for i := 1; i < len(stra); i++ {
		newnum, _ := strconv.Atoi(stra[i:len(stra)])
		_, presence := primes[newnum]
		if !presence {
			return false
		}
		newnum, _ = strconv.Atoi(stra[0 : len(stra)-i])
		_, presence = primes[newnum]
		if !presence {
			return false
		}
	}
	return true
}

func Euler37() string {
	sum := 0
	sieve := primes.FastSieve()
	count := 0
	primes := make(map[int]int)
	for prime := <-sieve; count < 11; prime = <-sieve {
		primes[prime] = prime
		if prime < 10 {
			continue
		}
		if isTruncable(prime, primes) {
			sum += prime
			count++
		}
	}
	return fmt.Sprint(sum)
}
