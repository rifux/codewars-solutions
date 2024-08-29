/*
	Count the number of divisors of a positive integer n.

	Random tests go up to n = 500000, but fixed tests go higher.

	Examples (input --> output)
	4 --> 3 // we have 3 divisors - 1, 2 and 4
	5 --> 2 // we have 2 divisors - 1 and 5
	12 --> 6 // we have 6 divisors - 1, 2, 3, 4, 6 and 12
	30 --> 8 // we have 8 divisors - 1, 2, 3, 5, 6, 10, 15 and 30
	Note you should only return a number, the count of divisors. The numbers between parentheses are shown only for you to see which numbers are counted in each case.

	(Source: https://www.codewars.com/kata/542c0f198e077084c0000c2e/train/go )
*/

package main

import (
	"fmt"
	"math"
)

type set map[int]struct{}

func (s set) add(item int) {
	s[item] = struct{}{}
}

func Divisors(n int) int {
	divisorsSet := set{}
	squareRoot := int(math.Sqrt(float64(n)))

	for i := 1; i <= squareRoot+1; i++ {
		if n%i == 0 {
			divisorsSet.add(i)
			divisorsSet.add(n / i)
		}
	}

	return len(divisorsSet)
}

func main() {
	fmt.Println(Divisors(1))
	fmt.Println(Divisors(10))
	fmt.Println(Divisors(11))
	fmt.Println(Divisors(54))
	fmt.Println(Divisors(1234567890000))
}
