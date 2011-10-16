package main

import (
	"projecteuler"
	"fmt"
	"profiling"
)

func main() {
	for i := 1; i < 47; i++ {
		if i == 10 {
			//problems takes forever, needs optimization
			continue
		}
		time, result := profiling.TimeArg(projecteuler.Euler, i)
		fmt.Printf("Problem %v : %v\t%v s\n", i, result, time/1000000000.0)
	}
}
