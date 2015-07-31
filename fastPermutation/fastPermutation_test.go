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
	//  50!-1
	k := 3.0414093e+64 - 1
	n := 50
	expected := make([]int, n)
	for i := 0; i < n; i++ {
		expected[i] = n - i
	}

	// Act
	result := FindPermut(n, k)

	// Assert
	assert.Equal(t, expected, result)
}

func Test_FindPermut_Should_ReturnRevertedSet_When_19(t *testing.T) {
	// Arrange
	n := 19
	t.Log("Running for ", n)
	k := 121645100408831999
	expected := make([]int, n)
	for i := 0; i < n; i++ {
		expected[i] = n - i
	}

	// Act
	result := FindPermut(n, float64(k))

	// Assert
	assert.Equal(t, expected, result)
}

func Test_FindPermut_Should_ReturnRevertedSet_When_18(t *testing.T) {
	// Arrange
	n := 18
	t.Log("Running for ", n)
	k := Factorial(n) - 1
	expected := make([]int, n)
	for i := 0; i < n; i++ {
		expected[i] = n - i
	}

	// Act
	result := FindPermut(n, float64(k))

	// Assert
	assert.Equal(t, expected, result)
}

func Test_FindPermut_Should_ReturnRevertedSet_When_MaxK(t *testing.T) {
	t.Skip("enable when bug for n=19 is fixed")
	// Arrange
	for n := 10; n < 30; n++ {
		t.Log("Running for ", n)
		k := Factorial(n) - 1
		expected := make([]int, n)
		for i := 0; i < n; i++ {
			expected[i] = n - i
		}

		// Act
		result := FindPermut(n, float64(k))

		// Assert
		assert.Equal(t, expected, result)
	}
}

func Test_Factorial(t *testing.T) {
	// Arrange
	expected := uint64(2432902008176640000)

	// Act
	result := Factorial(20)

	// Assert
	assert.Equal(t, expected, result)
}

func Test_FactorialMinusOne_Should_ReturnFactorialMinusOneForN(t *testing.T) {
	// Arrange
	expected := uint64(2432902008176639999)

	// Act
	result := FactorialMinusOne(20)

	// Assert
	assert.Equal(t, expected, result)
}

func Factorial(n int) uint64 {
	bigN := uint64(n)
	if n == 0 {
		return 1
	}

	result := uint64(1)

	for i := uint64(1); i <= bigN; i++ {
		result *= i
	}

	return result
}

func FactorialMinusOne(n int) uint64 {
	factorial := Factorial(n)

	return factorial - 1
}
