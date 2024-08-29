/*
	Given a string of digits, you should replace any digit below 5 with '0' and any digit 5 and above with '1'. Return the resulting string.

	Note: input will never be an empty string

	(Source: https://www.codewars.com/kata/57eae65a4321032ce000002d/train/go )
*/

package main

import (
	"fmt"
	"strconv"
)

func FakeBin(x string) (result string) {
	for _, digitRune := range x {
		digitInt, _ := strconv.Atoi(string(digitRune))
		if digitInt < 5 {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}

func main() {
	fmt.Println(FakeBin("15889923"))
}
