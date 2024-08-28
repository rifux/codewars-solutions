package main

import "fmt"

func Past(h, m, s int) int {
	return (h*3600 + m*60 + s) * 1000
}

func test(origin string, given int, expected int) {
	fmt.Printf("original data: %s\ngiven: %d; expected %d\nequals: %t\n\n", origin, given, expected, given == expected)
}

func main() {
	test("0,1,1", Past(0, 1, 1), 61000)
	test("1,1,1", Past(1, 1, 1), 3661000)
	test("0,0,0", Past(0, 0, 0), 0)
	test("1,0,1", Past(1, 0, 1), 3601000)
	test("1,0,0", Past(1, 0, 0), 3600000)
}
