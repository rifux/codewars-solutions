/*	https://www.codewars.com/kata/57eadb7ecd143f4c9c0000a3/train/go

	Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.

	The output should be two capital letters with a dot separating them.

	It should look like this:

	Sam Harris => S.H

	patrick feeney => P.F
*/

package main

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	words := strings.Split(name, " ")
	return fmt.Sprintf("%s.%s", strings.ToUpper(string(words[0][0])), strings.ToUpper(string(words[1][0])))
}
