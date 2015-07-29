package main

import (
	"fmt"
	"os"
	"permut/io"
	"permut/permutation"
	"strings"
)

func main() {
	err := io.ReadIntsLines(os.Stdin, Solve)
	if err != nil {
		fmt.Println(err)
	}
}

func Solve(n int, k int) {
	result := permutation.FindPermut(n, k)
	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
}
