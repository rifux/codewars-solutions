package main

import "fmt"

func MakeNegative(x int) int {
	if x > 0 {
		x *= -1
	}
	return x
}

func test(origin int, given int, expected int) {
	fmt.Printf("original data: %d\ngiven: %d; expected %d\nequals: %t\n\n", origin, given, expected, given == expected)
}

func main() {
	test(42, MakeNegative(42), -42)
	test(0, MakeNegative(0), 0)
	test(-10, MakeNegative(-10), -10)
}
