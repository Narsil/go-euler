//Problem 35
//How many circular primes are there below one million?

package projecteuler

import (
	"fmt"
	"big"
	"primes"
	"arithmetics"
	"strconv"
)

func Euler35() string {
	num := 0
	sieve := primes.FastSieve()
	circularPrimes := make(map[int]int)
	for a := <-sieve; a < 1000000; a = <-sieve {
		_, tested := circularPrimes[a]
		if tested {
			continue
		} else {
			//Test if it's circular
			circular := true
			primesTested := make(map[int]int)
			digits := []byte(strconv.Itoa(a))
			finishedRotations := false
			for ; ; arithmetics.Rotate(digits) {
				number, _ := strconv.Atoi(string(digits))
				_, finishedRotations = primesTested[number]
				if finishedRotations {
					break
				}
				primesTested[number] = 0
				if circular && !big.ProbablyPrime(big.NewInt(int64(number)), 1) {
					circular = false
				}
			}
			for n := range primesTested {
				if circular {
					num += 1
				}
				circularPrimes[n] = 0
			}
		}
	}
	return fmt.Sprint(num)
}
