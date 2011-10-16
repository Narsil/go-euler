//Problem 40
//Finding the nth digit of the fractional part of the irrational number.

package projecteuler

import (
	"fmt"
	"strconv"
	"bytes"
)

func Euler40() string {
	str := bytes.NewBufferString("")
	for i := 0; i < 100000; i++ {
		fmt.Fprint(str, i)
	}
	prod := 1
	power := 1
	for i := 0; i < 6; i++ {
		num, _ := strconv.Atoi(string(str.String()[power]))
		prod *= num
		power *= 10
	}
	return fmt.Sprint(prod)
}
