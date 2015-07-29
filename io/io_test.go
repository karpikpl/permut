package io

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_ReadInts_Should_ParseString(t *testing.T) {
	t.Log("ReadInts should parse string '32 67'")
	// Arrange
	testData := "32 67"
	expected := []int{32, 67}

	// Act
	result, err := ReadInts(strings.NewReader(testData))

	// Assert
	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}
