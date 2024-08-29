/*
	Given a non-empty array of integers, return the result of multiplying the values together in order. Example:

	[1, 2, 3, 4] => 1 * 2 * 3 * 4 = 24

	(Source: https://www.codewars.com/kata/57f780909f7e8e3183000078/train/go )
*/

package main

import "fmt"

func Grow(arr []int) (result int) {
	result = 1
	for _, num := range arr {
		result *= num
	}
	return result
}

func main() {
	fmt.Println(Grow([]int{1, 2, 3}))
}
