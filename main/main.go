package main

import (
	"fmt"
	"os"
	"permut/io"
	"permut/permutation"
	"strings"
)

func main() {
	io.ExecuteActionOnEachLine(os.Stdin, Solve)
}

func Solve(n int, k int) {
	result := permutation.FindPermut(n, k)
	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
}
