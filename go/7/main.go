/*
	Given the triangle of consecutive odd numbers:

					1
				3     5
			7     9    11
		13    15    17    19
	21    23    25    27    29
	...
	Calculate the sum of all the numbers of this triangle (starting at index 1) e.g.: (Input --> Output)

	1 -->  1
	2 --> 3 + 5 = 8
*/

package main

import "fmt"

func RowSumOddNumbers(n int) (result int) {
	result = n * (n + 1) / 2
	return result * result
}

/*
	EXPLANATION:

	1. Count items in rows:
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

	2. Calculate the sum of these items:
	- Every item can be calculated as the sum of the previous (n-1) item and 2 to keep
	  the numbers odd:
	    current_item = previous_item + 2
	  This means that the easiest way is to use a `for` loop, but if we want to
	  have an O(1), not O(n) difficulty solution, we will need to use this approach:
	    To find a formula for the sum of the series
		  1 + 3 + 5 + 7 + ...,
		we can first notice that this is an arithmetic series with a common
		difference of 2; the n'th term of an arithmetic series can be represented
		as
		  a_n = a_1 + (n-1)d,
		where a_1 is the first term and d is the common difference.

		In this case, a_1 = 1 and d = 2. So, the n'th term of the series is
		  a_n = 1 + (n-1)2 = 2n - 1.

		To find the sum of the first n terms of the series, we can use the
		formula for the sum of an arithmetic series:
		  S_n = n/2 * (a_1 + a_n).

		Substituting in the values of a_1 and a_n, we get:
		  S_n = n/2 * (1 + 2n - 1) = n/2 * (2n) = n^2.
*/

func main() {
	fmt.Println(RowSumOddNumbers(1))
	fmt.Println(RowSumOddNumbers(2))
	fmt.Println(RowSumOddNumbers(3))
	fmt.Println(RowSumOddNumbers(4))
}
