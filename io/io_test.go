package io

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_ExecuteActionOnEachLine_Should_CallPassedActionWithInts(t *testing.T) {
	// Arrange
	testData := "32 67\n10 15"
	var calledTimes int
	dummyAction := func(n int, k int) {
		calledTimes++
	}

	// Act
	ExecuteActionOnEachLine(strings.NewReader(testData), dummyAction)

	// Assert
	assert.Equal(t, 2, calledTimes)
}
