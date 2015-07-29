package permutation

import (
	"fmt"
)

func FindPermut(n int, k int) []int {
	table := make([]int, n)
	for i := 1; i <= n; i++ {
		table[i-1] = i
	}

	for i := 0; i < k; i++ {
		PermutMe(&table)
	}

	return table
}

func PermutMe(tablePtr *[]int) {
	table := *tablePtr

	// Find the largest index k such that a[k] < a[k + 1]. If no such index exists, the permutation is the last permutation.
	k := FindK(tablePtr)

	if k < 0 {
		fmt.Println("Last permutation found", table)
		return
	}

	// Find the largest index l greater than k such that a[k] < a[l].
	l := FindL(tablePtr, k)

	if l < 0 {
		panic("l shouldn't be less than 0")
	}

	// Swap the value of a[k] with that of a[l].
	tmp := table[k]
	table[k] = table[l]
	table[l] = tmp

	// Reverse the sequence from a[k + 1] up to and including the final element a[n].
	Reverse(&table, k+1)
}

func Reverse(sPtr *[]int, from int) {
	s := *sPtr
	for i, j := from, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func FindK(sPtr *[]int) int {
	// Find the largest index k such that a[k] < a[k + 1]. If no such index exists, the permutation is the last permutation.
	table := *sPtr
	for index := len(table) - 2; index >= 0; index-- {
		if table[index] < table[index+1] {
			return index
		}
	}
	return -1
}

func FindL(sPtr *[]int, k int) int {
	// Find the largest index l greater than k such that a[k] < a[l].
	table := *sPtr
	for index := len(table) - 1; index > k; index-- {
		if table[index] > table[k] {
			return index
		}
	}
	return -1
}
