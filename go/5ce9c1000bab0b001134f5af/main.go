package main

import "fmt"

func QuarterOf(month int) (result int) {
	result = month / 3
	if month%3 != 0 {
		result += 1
	}
	return result
}

func test(origin int, given int, expected int) {
	fmt.Printf("original data: %d\ngiven: %d; expected %d\nequals: %t\n\n", origin, given, expected, given == expected)
}

func main() {
	test(3, QuarterOf(3), 1)
	test(8, QuarterOf(8), 3)
	test(11, QuarterOf(11), 4)
}
