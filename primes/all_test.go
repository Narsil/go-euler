package primes

import (
	"testing"
	"fmt"
)

var primes = []uint64{2, 3, 5, 7, 11, 13, 17, 19}

func TestSieve(t *testing.T) {
	sieve := Sieve()
	for _, prime := range primes {
		p := <-sieve
		if prime != p {
			t.Errorf("Sieve is %v, want %v\n", p, prime)
		}
	}
}

func TestFactors(t *testing.T) {
	x := uint64(100)
	fmt.Printf("%v", Factors(x))
}
