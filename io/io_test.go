package io

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"strings"
	"testing"
)

func Test_ExecuteActionOnEachLine_Should_CallPassedActionWManyTimes(t *testing.T) {
	// Arrange
	testData := "32 67\n50 30414093201713378043612608166064768844377641568960511999999999999"
	var calledTimes int
	dummyAction := func(n int, k *big.Int) {
		calledTimes++
		t.Log("read k=", k)
	}

	// Act
	ExecuteActionOnEachLine(strings.NewReader(testData), dummyAction)

	// Assert
	assert.Equal(t, 2, calledTimes)
}

func Test_ExecuteActionOnEachLine_Should_CallPassedActionCorrectValues(t *testing.T) {
	// Arrange
	bigNumber := "30414093201713378043612608166064768844377641568960511999999999999"
	expected, _ := new(big.Int).SetString(bigNumber, 10)
	testData := "50 " + bigNumber
	var result *big.Int
	dummyAction := func(n int, k *big.Int) {
		result = k
		t.Log("read k=", k)
	}

	// Act
	ExecuteActionOnEachLine(strings.NewReader(testData), dummyAction)

	// Assert
	assert.Equal(t, expected, result)
}
