package main

import "fmt"

func test(origin string, given int, expected int) {
	fmt.Printf("original data: %s\ngiven: %d; expected %d\nequals: %t\n\n", origin, given, expected, given == expected)
}

func Move(position int, roll int) int {
	return position + roll*2
}

func main() {
	test("0,4", Move(0, 4), 8)
	test("3,6", Move(3, 6), 15)
	test("2,5", Move(2, 5), 12)
}
