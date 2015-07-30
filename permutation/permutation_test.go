package permutation

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func Test_Reverse_Should_ReverseArrayStartingFromK(t *testing.T) {
	// Arrange
	array := []int{1, 3, 5, 2, 4}
	expected := []int{1, 3, 4, 2, 5}
	from := 2

	// Act
	Reverse(&array, from)

	// Assert
	assert.Equal(t, expected, array, "reverse shuld return array with reversed elements starting from k'th")
}

func Test_Reverse_Should_ReverseArray(t *testing.T) {
	// Arrange
	array := []int{1, 3, 5, 2, 4}
	expected := []int{4, 2, 5, 3, 1}

	// Act
	Reverse(&array, 0)

	// Assert
	assert.Equal(t, expected, array, "reverse shuld return array with reversed elements starting from k'th")
}

func Test_PermutMe_Should_PermutateTheTable(t *testing.T) {
	// Arrange
	array := []int{2, 1, 4, 3}
	expected := []int{2, 3, 1, 4}

	// Act
	PermutMe(&array)

	// Assert
	assert.Equal(t, expected, array)
}

func Test_FindK_Should_ReturnK(t *testing.T) {
	// Arrange
	// Find the largest index k such that a[k] < a[k + 1]. If no such index exists, the permutation is the last permutation.
	array := []int{3, 4, 2, 1}

	// Act
	result := FindK(&array)

	// Assert
	assert.Equal(t, 0, result)
}

func Test_FindK_Should_ReturnK2(t *testing.T) {
	// Arrange
	// Find the largest index k such that a[k] < a[k + 1]. If no such index exists, the permutation is the last permutation.
	array := []int{1, 2, 3, 4}

	// Act
	result := FindK(&array)

	// Assert
	assert.Equal(t, 2, result)
}

func Test_FindK_Should_ReturnL(t *testing.T) {
	// Arrange
	// Find the largest index l greater than k such that a[k] < a[l].
	array := []int{1, 2, 3, 4}
	k := 2

	// Act
	result := FindL(&array, k)

	// Assert
	assert.Equal(t, 3, result)
}

func Test_FindPermut_Should_Find19PermutationOfSet10(t *testing.T) {
	t.Log("Finds 4th permutation of 3 size set")
	// Arrange
	expected := []int{1, 2, 3, 4, 5, 6, 7, 9, 10, 8}
	k := 3
	n := 10

	// Act
	result := FindPermut(n, k)

	// Assert
	assert.Equal(t, expected, result, "3'rd permutation should be equal to 1 2 3 4 5 6 7 9 10 8")
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

func BenchmarkFindPermut(b *testing.B) {
	//big := 3.0414093e+64
	big := 1<<31 - 1

	for i := 0; i < b.N; i++ {
		FindPermut(50, big)
	}
}
