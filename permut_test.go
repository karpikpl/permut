package main

import (
  "testing"
  "fmt"
  "github.com/stretchr/testify/assert"
  "strings"
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
  expected := [3]int{3,1,2}
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
  expected := [4]int{4,1,3,2}
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

func BenchmarkHello(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }
}
