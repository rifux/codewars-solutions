/*
	Given the triangle of consecutive odd numbers:

					1
				3     5				4
			7     9    11			6
		13    15    17    19		8
	21    23    25    27    29		10
31    33    35    37    39    41	12
	...
	Calculate the sum of the numbers in the nth row of this triangle (starting at index 1) e.g.: (Input --> Output)

	1 -->  1
	2 --> 3 + 5 = 8
*/

package main

import "fmt"

func RowSumOddNumbers(n int) int {
	/*
		items := n * (n - 1) / 2
		a, b := 2*(items+1)-1, 2*(items+n)-1
		return n * (b + a) / 2
	*/
	return n * n * n
}

func main() {
	fmt.Println(RowSumOddNumbers(2))
}

/*
	1. Find the number of items in the n'th row:
	- Every next row we are seeing only one new item (number), so
	  the number of items in row should be same as the row's number:
	    items := n

	2. Count items in rows:
	- For every new row, there is one new element, which means that the number of items
	  on the n'th row can be calculated using the formula:
		items = n
	  We also need to calculate the sum of all previous rows and the current row
	  using the formula:
		items_sum = n + (n-1) + (n-2) + ...
	  which equals to
		items_sum = 1 + 2 + 3 + ... + (n-2) + (n-1) + n
	  or
		items_sum = n * (n + 1) / 2
	  (Source: https://en.wikipedia.org/wiki/1_%2B_2_%2B_3_%2B_4_%2B_%E2%8B%AF )

	3. Source: https://en.wikipedia.org/wiki/Arithmetic_progression#Sum

	4. Simplify
*/
