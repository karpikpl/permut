package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_ReadInts_Should_ParseString(t *testing.T) {
	t.Log("ReadInts should parse string '32 67'")
	// Arrange
	testData := "32 67"

	// Act
	result, err := ReadInts(strings.NewReader(testData))

	// Assert
	assert.Equal(t, 32, result[0])
	assert.Equal(t, 67, result[1])
	assert.Nil(t, err)
}

func Test_FindPermut_Should_Find4PermutationOfSet3(t *testing.T) {
	t.Log("Finds 4th permutation of 3 size set")
	// Arrange
	expected := []int{3, 1, 2}
	k := 4
	n := 3

	// Act
	result := FindPermut(n, k)

	// Assert
	assert.Equal(t, expected, result, "k'th permutation should be equal to 3 1 2")
}

func Test_FindPermut_Should_Find19PermutationOfSet4(t *testing.T) {
	t.Log("Finds 4th permutation of 3 size set")
	// Arrange
	expected := []int{4, 1, 3, 2}
	k := 19
	n := 4

	// Act
	result := FindPermut(n, k)

	// Assert
	assert.Equal(t, expected, result, "k'th permutation should be equal to 4 1 3 2")
}

func Test_Factorial_Should_Be120When5(t *testing.T) {
	// Arrange
	expected := 120
	n := 5

	// Act
	result := Factorial(n)

	// Assert
	assert.Equal(t, expected, result, "factorial of 6 should be 120")
}

func Test_Reverse_Should_ReverseArrayStartingFromK(t *testing.T) {
	// Arrange
	array := []int{1, 3, 5, 2, 4}
	expected := []int{1, 3, 4, 2, 5}
	from := 2

	// Act
	result := Reverse(array, from)

	// Assert
	assert.Equal(t, expected, result, "reverse shuld return array with reversed elements starting from k'th")
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
