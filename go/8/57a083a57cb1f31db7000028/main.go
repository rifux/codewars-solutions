/*
	Complete the function that takes a non-negative integer n as input, and returns a list of all the powers of 2 with the exponent ranging from 0 to n ( inclusive ).
*/

package main

import "fmt"

func PowersOfTwo(n int) (result []uint64) {
	result = append(result, 1)
	for i := 0; i < n; i++ {
		result = append(result, result[i]*2)
	}
	return result
}

func main() {
	fmt.Println(PowersOfTwo(4))
}
