package main

import (
	"fmt"
	"os"
	"permut/io"
	"permut/fastPermutation"
	"strings"
)

func main() {
	io.ExecuteActionOnEachLine(os.Stdin, Solve)
}

func Solve(n int, k int) {
	result := fastPermutation.FindPermut(n, k)
	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
}
