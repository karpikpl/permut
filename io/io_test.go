package io

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_ExecuteActionOnEachLine_Should_CallPassedActionWithInts(t *testing.T) {
	// Arrange
	testData := "32 67\n50 3.0414093e+64"
	var calledTimes int
	dummyAction := func(n int, k float64) {
		calledTimes++
	}

	// Act
	ExecuteActionOnEachLine(strings.NewReader(testData), dummyAction)

	// Assert
	assert.Equal(t, 2, calledTimes)
}
