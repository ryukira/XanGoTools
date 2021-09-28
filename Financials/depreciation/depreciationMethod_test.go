package depreciation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStraightLine(t *testing.T) {
	result := StraightLine(2000, 5, 500)
	assert.EqualValues(t, 500, result)
}

func TestStraightLineArr(t *testing.T) {
	result := StraightLineArr(2000, 5, 500)
	expected := make([]float64, 0)
	expected = append(expected, 2000)
	expected = append(expected, 1700)
	expected = append(expected, 1400)
	expected = append(expected, 1100)
	expected = append(expected, 800)
	assert.EqualValues(t, expected, result)
}
