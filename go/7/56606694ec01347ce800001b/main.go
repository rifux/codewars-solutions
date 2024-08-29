/*
	Implement a function that accepts 3 integer values a, b, c. The function should return true if a triangle can be built with the sides of given length and false in any other case.

	(In this case, all triangles must have surface greater than 0 to be accepted).

	Examples:

	Input -> Output
	1,2,2 -> true
	4,2,3 -> true
	2,2,2 -> true
	1,2,3 -> false
	-5,1,3 -> false
	0,2,3 -> false
	1,2,9 -> false

	(Source: https://www.codewars.com/kata/56606694ec01347ce800001b/train/go )
*/

package main

import (
	"fmt"
	"sort"
)

func IsTriangle(a, b, c int) bool {
	sides := []int{a, b, c}
	sort.Ints(sides)
	if sides[2] >= sides[1]+sides[0] {
		return false
	}
	return true
}

func main() {
	fmt.Println(IsTriangle(1, 2, 2))
	fmt.Println(IsTriangle(4, 2, 3))
	fmt.Println(IsTriangle(2, 2, 2))
	fmt.Println(IsTriangle(1, 2, 3))
	fmt.Println(IsTriangle(5, 1, 2))
}
