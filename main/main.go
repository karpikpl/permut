package main

import (
	"fmt"
	"os"
	"permut/fastPermutation"
	"permut/io"
	"strings"
)

func main() {
	io.ExecuteActionOnEachLine(os.Stdin, Solve)
}

func Solve(n int, k float64) {
	result := fastPermutation.FindPermut(n, k)
	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
}
