//Problem 22
// What is the total of all the name scores in the file

package projecteuler

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sort"
)

func value(s string) int {
	sum := 0
	for _, char := range s {
		sum += char - 64
	}
	return sum
}

func Euler22() string {
	sum := 0
	file, error := ioutil.ReadFile("names.txt")
	if error != nil {
		return "Cannot open names.txt, make sure you have it in the current directory" + error.String()
	}
	names := strings.SplitN(string(file), ",", -1)
	sort.Strings(names)
	for index, name := range names {
		sum += value(strings.Trim(name, "\"")) * (index + 1)
	}
	return fmt.Sprint(sum)
}
