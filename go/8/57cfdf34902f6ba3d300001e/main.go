package main

import (
	"fmt"
	"sort"
)

func addStars(original string) (result string) {
	for _, letter := range original {
		result += string(letter) + "***"
	}
	return result[:len(result)-3]
}

func TwoSort(arr []string) string {
	sort.Strings(arr)
	return addStars(arr[0])
}

func Expect(given string, shouldBe string) {
	fmt.Printf("%s equals %s:  %t\n", given, shouldBe, given == shouldBe)
}

func main() {
	var s []string
	s = []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}
	Expect(TwoSort(s), "b***i***t***c***o***i***n")
	s = []string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}
	Expect(TwoSort(s), "a***r***e")
	s = []string{"lets", "talk", "about", "javascript", "the", "best", "language"}
	Expect(TwoSort(s), "a***b***o***u***t")
	s = []string{"i", "want", "to", "travel", "the", "world", "writing", "code", "one", "day"}
	Expect(TwoSort(s), "c***o***d***e")
	s = []string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}
	Expect(TwoSort(s), "L***e***t***s")
}
