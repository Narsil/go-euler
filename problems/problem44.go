//Problem 44
//Find the smallest pair of pentagonal numbers whose sum and difference is pentagonal.

package projecteuler

import (
	"fmt"
	"math"
)

func isPentagonal(n uint64) bool {
	sq := math.Sqrt(float64(1 + 24*n))
	return sq == float64(uint64(sq)) && (1+uint64(sq))%6 == 0
}

func Euler44() string {
	pentagonals := make(map[uint64]uint64)
	for i := uint64(1); ; i++ {
		pent1 := i * (uint64(3)*i - uint64(1)) / uint64(2)
		for pent2 := range pentagonals {
			diff := pent1 - pent2
			_, presence := pentagonals[diff]
			if presence {
				sum := pent1 + pent2
				if isPentagonal(sum) {
					return fmt.Sprint(diff)
				}
			}
		}
		pentagonals[pent1] = 1
	}
	return "Error occured"
}
