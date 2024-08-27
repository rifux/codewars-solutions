package katatest

import (
	. "codewarrior/kata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var arr2 = []bool{
	true, true, true, false,
	true, true, true, true,
	true, false, true, false,
	true, false, false, true,
	true, true, true, true,
	false, false, true, true,
}

var Test = CountSheeps(arr2) == 17
var _ = Describe("Test Example", func() {
	It("There are 17 sheeps in total", func() {
		arr1 := []bool{
			true, true, true, false,
			true, true, true, true,
			true, false, true, false,
			true, false, false, true,
			true, true, true, true,
			false, false, true, true,
		}
		Expect(CountSheeps(arr1)).To(Equal(17))
	})
})
