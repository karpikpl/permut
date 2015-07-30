package fastPermutation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Benchmark_FindPermut(b *testing.B) {
	big := float64(1<<31 - 1)
	for i := 0; i < b.N; i++ {
		FindPermut(50, big)
	}
}

func Test_FindPermut_Should_Find4PermutationOfSet3(t *testing.T) {
	t.Log("Finds 4th permutation of 3 size set")
	// Arrange
	expected := []int{3, 1, 2}
	k := float64(4)
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
	k := float64(19)
	n := 4

	// Act
	result := FindPermut(n, k)

	// Assert
	assert.Equal(t, expected, result, "k'th permutation should be equal to 4 1 3 2")
}

func Test_FindPermut_Should_WorkForBigNumbers(t *testing.T) {
	// Arrange
	k := 3.0414093e+64 - 1
	n := 50
	expected := make([]int, n)
	for i := 0; i < n; i++ {
		expected[i] = 50 - i
	}

	// Act
	result := FindPermut(n, k)

	// Assert
	assert.Equal(t, expected, result)
}
