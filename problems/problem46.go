//Problem 46
//What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?

package projecteuler

import (
	"fmt"
	"primes"
	"math"
)

func isSquare(a int) bool {
	sq := math.Sqrt(float64(a))
	return sq == float64(int(sq))
}

func Euler46() string {
	sieve := primes.FastSieve()
	primes := make(map[int]int)
	for i := 0; i < 10000; i++ {
		prime := <-sieve
		primes[i] = prime
	}
	for compo := 9; ; compo += 2 {
		found := true
		for i := 0; i < compo; i++ {
			diff := compo - primes[i]
			if diff%2 == 0 && isSquare((diff)/2) {
				found = false
				break
			}
		}
		if found {
			return fmt.Sprint(compo)
		}
	}
	return "Problem occured"
}
