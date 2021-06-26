package interest

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var interestRate float64 = 0.10 // as 10%
var timeRate float64 = 10
var valueRate float64 = 1 // like $1

func LogTestValue(result float64, expected float64, t *testing.T) {
	var isUnderControl bool
	isUnderControl = false
	if (result/expected) > 0.8 && (result/expected) < 1.2 {
		isUnderControl = true
	}
	fmt.Printf("counted is %f and expected is %f \n", result, expected)
	fmt.Printf("the difference is %f \n", result/expected)
	assert.EqualValues(t, true, isUnderControl)
}

func TestFPCalculation(t *testing.T) {
	var expected, result float64
	expected = 2.594
	result = FuturePresent(valueRate, timeRate, interestRate)
	LogTestValue(result, expected, t)
}

func TestPFCalculation(t *testing.T) {
	var expected, result float64
	expected = 0.3855
	result = PresentFuture(valueRate, timeRate, interestRate)
	LogTestValue(result, expected, t)
}

func TestFACalculation(t *testing.T) {
	var expected, result float64
	expected = 15.937
	result = FutureAnnual(valueRate, timeRate, interestRate)
	LogTestValue(result, expected, t)
}

func TestAFCalculation(t *testing.T) {
	var expected, result float64
	expected = 0.0627
	result = AnnualFuture(valueRate, timeRate, interestRate)
	LogTestValue(result, expected, t)
}

func TestAPCalculation(t *testing.T) {
	var expected, result float64
	expected = 0.1627
	result = AnnualPresent(valueRate, timeRate, interestRate)
	LogTestValue(result, expected, t)
}

func TestPACalculation(t *testing.T) {
	var expected, result float64
	expected = 6.145
	result = PresentAnnual(valueRate, timeRate, interestRate)
	LogTestValue(result, expected, t)
}

func TestAGCalculation(t *testing.T) {
	var expected, result float64
	expected = 3.725
	result = AnnualGradient(valueRate, timeRate, interestRate)
	LogTestValue(result, expected, t)
}

func TestPGCalculation(t *testing.T) {
	var expected, result float64
	expected = 22.891
	result = PresentGradient(valueRate, timeRate, interestRate)
	LogTestValue(result, expected, t)
}
