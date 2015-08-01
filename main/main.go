package main

import (
	"fmt"
	"math/big"
	"os"
	"permut/fastPermutation"
	"permut/io"
	"strings"
)

func main() {
	io.ExecuteActionOnEachLine(os.Stdin, Solve)
}

func Solve(n int, k *big.Int) {
	result := fastPermutation.FindPermut(n, k)
	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
}
