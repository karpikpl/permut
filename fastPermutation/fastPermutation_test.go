package fastPermutation

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_FindPermut(b *testing.B) {
	big, _ := new(big.Int).SetString("30414093201713378043612608166064768844377641568960511999999999999", 10)
	for i := 0; i < b.N; i++ {
		FindPermut(50, big)
	}
}

func Test_FindPermut_Should_Find4PermutationOfSet3(t *testing.T) {
	t.Log("Finds 4th permutation of 3 size set")
	// Arrange
	expected := []int{3, 1, 2}
	k := big.NewInt(4)
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
	k := big.NewInt(19)
	n := 4

	// Act
	result := FindPermut(n, k)

	// Assert
	assert.Equal(t, expected, result, "k'th permutation should be equal to 4 1 3 2")
}

func Test_FindPermut_Should_Work_When_N50(t *testing.T) {
	// Arrange
	//  50!-1
	k, _ := new(big.Int).SetString("30414093201713378043612608166064768844377641568960511999999999999", 10)
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

func Test_FindPermut_Should_ReturnRevertedSet_When_BigN(t *testing.T) {
	// Arrange
	for n := 10; n < 30; n++ {
		t.Log("Running for ", n)
		k := new(big.Int).Sub(Factorial(n), big.NewInt(1))
		expected := make([]int, n)
		for i := 0; i < n; i++ {
			expected[i] = n - i
		}

		// Act
		result := FindPermut(n, k)

		// Assert
		assert.Equal(t, expected, result)
	}
}

func Test_Factorial_Should_Work_When_N50(t *testing.T) {
	// Arrange
	expected, _ := new(big.Int).SetString("30414093201713378043612608166064768844377641568960512000000000000", 10)

	// Act
	result := Factorial(50)

	// Assert
	assert.Equal(t, expected, result)
}

func Test_FactorialMinusOne_Should_ReturnFactorialMinusOne_When_N50(t *testing.T) {
	// Arrange
	expected, _ := new(big.Int).SetString("30414093201713378043612608166064768844377641568960511999999999999", 10)

	// Act
	result := new(big.Int).Sub(Factorial(50), big.NewInt(1))

	// Assert
	assert.Equal(t, expected, result)
}

func Factorial(n int) *big.Int {
	bigN := int64(n)
	if n == 0 {
		return big.NewInt(1)
	}

	result := big.NewInt(1)

	for i := int64(1); i <= bigN; i++ {
		result = result.Mul(result, big.NewInt(i))
	}

	return result
}
