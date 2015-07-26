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

func BenchmarkHello(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }
}
