package io

import (
	"bufio"
	"io"
	"math/big"
	"strconv"
	"strings"
)

func ExecuteActionOnEachLine(r io.Reader, action func(int, *big.Int)) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")

		if len(values) == 2 {
			n, _ := strconv.Atoi(values[0])
			k, _ := new(big.Int).SetString(values[1], 10)
			action(n, k)
		} else {
			panic("something went wrong - do data to process")
		}
	}
}
