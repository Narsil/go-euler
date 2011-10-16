// Problem 16 - http://projecteuler.net/
//
// What is the sum of the digits of the number 2^1000?


package projecteuler

import (
	"fmt"
	"big"
	"bigextension"
	"strconv"
)

func Euler16() string {
	str := bigextension.Pow(big.NewInt(2), 1000).String()
	sum := 0
	digit := 0
	for _, char := range str {
		digit, _ = strconv.Atoi(string(char))
		sum += digit
	}
	return fmt.Sprint(sum)
}
