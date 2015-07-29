package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func Factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}

func FindPermut2(n int, k int) []int {
	// find k'th element of n-size permutation

	// there are n! permutations
	// each number is on first position (n-1)! times
	totalP := Factorial(n)
	// 19 -> 20
	// 4! -> each can be 6 times
	// 20/6 -> 3.x -> 4
	// 6 / 3 ->  2
	result := make([]int, n)
	result[0] = int(math.Ceil(float64(k+1.0) / (float64(totalP) / float64(n))))
	result[1] = totalP
	return result
}

func FindPermut(n int, k int) []int {
	table := make([]int, n)
	for i := 1; i <= n; i++ {
		table[i-1] = i
	}

	tableSlice := table[:]
	fmt.Printf("table is %T and tableSlice is %T and equal %d \n", table, tableSlice, tableSlice)

	for i := 0; i < k; i++ {
		tableSlice = PermutMe(tableSlice)
		fmt.Println("permutation", i+1, "=", tableSlice)
	}

	return table
}

func PermutMe(table []int) []int {
	var k, l int
	// Find the largest index k such that a[k] < a[k + 1]. If no such index exists, the permutation is the last permutation.
	// Find the largest index l greater than k such that a[k] < a[l].
	// Swap the value of a[k] with that of a[l].
	// Reverse the sequence from a[k + 1] up to and including the final element a[n].
	for index := len(table) - 1; index > 0; index-- {
		if table[index] > table[index-1] {
			k = index - 1
			break
		} else {
			fmt.Println("Last permutation found")
			return table
		}
	}
	//fmt.Println("Found k!",k)

	for index := len(table) - 1; index > k; index-- {
		if table[index] > table[k] {
			l = index
			break
		}
	}
	//fmt.Println("Found l!",l)

	tmp := table[k]
	table[k] = table[l]
	table[l] = tmp

	table = Reverse(table, k+1)

	return table
}

func Reverse(s []int, from int) []int {
	for i, j := from, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func main() {
	argsWithoutProg := os.Args[1:]
	var input []int
	var err error

	if len(argsWithoutProg) == 0 {
		input, err = ReadInts(os.Stdin)
	} else {
		input, err = ReadInts(strings.NewReader(argsWithoutProg[0]))
	}

	fmt.Println(input, err)
}
