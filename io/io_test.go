package io

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_ExecuteActionOnEachLine_Should_CallPassedActionWithInts(t *testing.T) {
	// Arrange
	testData := "32 67\n50 30414093201713378043612608166064768844377641568960511999999999999"
	var calledTimes int
	dummyAction := func(n int, k float64) {
		calledTimes++
		t.Log("read k=", k)
	}

	// Act
	ExecuteActionOnEachLine(strings.NewReader(testData), dummyAction)

	// Assert
	assert.Equal(t, 2, calledTimes)
}
