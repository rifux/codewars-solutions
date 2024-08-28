package main

import "fmt"

func CountBy(x, n int) (returnList []int) {
	returnList = append(returnList, x)
	for ; n > 1; n-- {
		returnList = append(returnList, returnList[len(returnList)-1]+x)
	}
	return returnList
}

func main() {
	fmt.Println(CountBy(1, 10))
	fmt.Println(CountBy(2, 5))
}
