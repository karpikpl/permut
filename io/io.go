package io

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ExecuteActionOnEachLine(r io.Reader, action func(int, float64)) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")

		if len(values) == 2 {
			n, _ := strconv.Atoi(values[0])
			k, _ := strconv.ParseFloat(values[1], 64)
			action(n, k)
		} else {
			panic("something went wrong - do data to process")
		}
	}
}
