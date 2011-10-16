//Problem 45
//After 40755, what is the next triangle number that is also pentagonal and hexagonal?

package projecteuler

import (
	"fmt"
	"math"
)

func isTriangle(n uint64) bool {
	sq := math.Sqrt(float64(1 + 8*n))
	return sq == float64(uint64(sq)) && (uint64(sq)-1)%2 == 0
}

func hexagonal(i uint64) uint64 {
	return i * (2*i - 1)
}

func Euler45() string {
	h := hexagonal(144)
	for i := uint64(144); !(isPentagonal(h) && isTriangle(h)); i++ {
		h = hexagonal(i)
	}
	return fmt.Sprint(h)
}
