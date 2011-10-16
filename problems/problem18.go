//Problem18
//Find the maximum sum travelling from the top of the triangle to the base.

package projecteuler

import (
	"fmt"
	"strings"
	"strconv"
)

func Euler18() string {
	str := "75\n" +
		"95 64\n" +
		"17 47 82\n" +
		"18 35 87 10\n" +
		"20 04 82 47 65\n" +
		"19 01 23 75 03 34\n" +
		"88 02 77 73 07 63 67\n" +
		"99 65 04 28 06 16 70 92\n" +
		"41 41 26 56 83 40 80 70 33\n" +
		"41 48 72 33 47 32 37 16 94 29\n" +
		"53 71 44 65 25 43 91 52 97 51 14\n" +
		"70 11 33 28 77 73 17 78 39 68 17 57\n" +
		"91 71 52 38 17 14 91 43 58 50 27 29 48\n" +
		"63 66 04 68 89 53 67 30 73 16 69 87 40 31\n" +
		"04 62 98 27 23 09 70 98 73 93 38 53 60 04 23"

		//Convert this in triangular matrix
	lines := strings.Split(str, "\n")
	digits := make([]string, 225)
	numbers := make([]int, 225)
	for i, line := range lines {
		digits = strings.SplitN(line, " ", -1)
		for j, digit := range digits {
			numbers[i*15+j], _ = strconv.Atoi(string(digit))
		}
	}
	//Starting from the bottom, add to the upper layer the best solution
	//This does not reveal the actual best path but does give the answer
	for i := 13; i >= 0; i-- {
		for j := 0; j < 15; j++ {
			if j > i {
				break
			}
			numbers[i*15+j] += maximum(numbers[(i+1)*15+j], numbers[(i+1)*15+j+1])
		}
	}
	return fmt.Sprint(numbers[0])
}
