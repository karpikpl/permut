package main

import (
	"fmt"
	"os"
	"permut/io"
	"permut/permutation"
)

func main() {
	for {
		input, err := io.ReadIntsLine(os.Stdin)

		if len(input) > 0 {
			result := permutation.FindPermut(input[0], input[1])
			fmt.Printf("%d\n", result)
		} else {
			fmt.Println(err)
			break
		}
	}
}
