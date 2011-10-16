//Problem 19
//How many Sundays fell on the first of the month during the twentieth century?

package projecteuler

import (
	"fmt"
)

func monthSize(month, year int) int {
	switch month {
	case 2:
		if year%4 != 0 {
			return 28
		} else if year%100 != 0 {
			return 29
		} else if year%400 != 0 {
			return 28
		} else {
			return 29
		}
		break
	case 4, 6, 9, 11:
		return 30
		break
	default:
		return 31
		break
	}
	return 0

}

func Euler19() string {
	year := 1901
	month := 1
	num := 0
	day := 6
	//Date of the first sunday of the century, we'll loop over all sundays
	//increment  when day==1
	for ; year < 2001; day += 7 {
		MonthSize := monthSize(month, year)
		if day > MonthSize {
			month++
			if month > 12 {
				month = 1
				year++
			}
			day -= MonthSize
		}
		if day == 1 {
			num++
			continue
		}
	}
	return fmt.Sprint(num)
}
