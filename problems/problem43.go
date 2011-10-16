//Problem 43
//Find the sum of all pandigital numbers with an unusual sub-string divisibility property.

package projecteuler

import (
	"fmt"
	"arithmetics"
	"strconv"
)

var smallprimes = []uint64{2, 3, 5, 7, 11, 13, 17}

func Euler43() string {
	sum := uint64(0)
	pand := []byte("1023456789")
	for arithmetics.Permute(pand) {
		hasProperty := true
		for i, prime := range smallprimes {
			num, _ := strconv.Atoui64(string(pand[1+i : 4+i]))
			if num%prime != 0 {
				hasProperty = false
				break
			}
		}
		if hasProperty {
			num, _ := strconv.Atoui64(string(pand))
			sum += num
		}
	}
	return fmt.Sprint(sum)
}
