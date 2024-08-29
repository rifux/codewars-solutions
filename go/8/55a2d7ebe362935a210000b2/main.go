/*
	Given an array of integers your solution should find the smallest integer.

	For example:

	Given [34, 15, 88, 2] your solution will return 2
	Given [34, -345, -1, 100] your solution will return -345
	You can assume, for the purpose of this kata, that the supplied array will not be empty.

	(Source: https://www.codewars.com/kata/55a2d7ebe362935a210000b2/train/go )
*/

package main

import (
	"fmt"
	"sort"
)

func SmallestIntegerFinder(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}

func main() {
	fmt.Println(SmallestIntegerFinder([]int{34, 15, 88, 2}))
	fmt.Println(SmallestIntegerFinder([]int{34, -345, -1, 100}))
}
