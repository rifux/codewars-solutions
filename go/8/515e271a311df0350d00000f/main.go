/*
	Complete the square sum function so that it squares each number passed into it and then sums the results together.

	For example, for [1, 2, 2] it should return 9 because 1^2 + 2^2 + 2^2 = 9.
*/

package main

import "fmt"

func SquareSum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number * number
	}
	return result
}

func main() {
	fmt.Println(SquareSum([]int{1, 2, 2}))
	fmt.Println(SquareSum([]int{0, 3, 4, 5}))
	fmt.Println(SquareSum([]int{}))
}
