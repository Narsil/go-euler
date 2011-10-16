//Problem 20
// Find the sum of the digits of 100!

package projecteuler

import (
	"fmt"
	"big"
	"strconv"
)

func Euler20() string {
	n := big.NewInt(0)
	str := n.MulRange(1, 100).String()
	num := 0
	digit := 0
	for _, char := range str {
		digit, _ = strconv.Atoi(string(char))
		num += digit
	}
	return fmt.Sprint(num)
}
