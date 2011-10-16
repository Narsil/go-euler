//Problem39
//If p is the perimeter of a right angle triangle, {a, b, c}, which value, for p ≤ 1000, has the most solutions?

package projecteuler

import (
	"fmt"
	"math"
)

func Euler39() string {
	pythagorean := make(map[int]int)
	for i := 0; i < 1000; i++ {
		for j := i + 1; j < 1000; j++ {
			k := math.Sqrt(float64(i*i + j*j))
			if k == float64(int64(k)) && i+j+int(k) < 1000 {
				pythagorean[i+j+int(k)]++
			}
		}
	}
	max := 0
	maxPerimeter := 0
	for perimeter, hits := range pythagorean {
		if hits > max {
			max = hits
			maxPerimeter = perimeter
		}
	}
	return fmt.Sprint(maxPerimeter)
}
