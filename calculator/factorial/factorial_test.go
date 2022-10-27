package factorial

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testedValues struct {
	name     string
	input    int
	expected int
}

func TestFactorial(t *testing.T) {
	sliceOfStruct := []testedValues{
		{"test with 0", 0, 1},
		{"test with 1", 1, 1},
		{"test with 2", 2, 2},
		{"test with 5", 5, 120},
		{"test with 8", 8, 40320},
		{"test with 10", 10, 3628800},
		{"test with 12", 12, 479001600},
	}

	for _, val := range sliceOfStruct {
		/*
			val := val
			t.Run(val.name, func(t *testing.T) {
				res := Factorial(val.input)

				if res != val.expected {
					t.Errorf("For %d expexted %d, got %d", val.input, val.expected, res)
				}
			})
			Вариант многопоточной реализации T-теста без testify.
		*/
		res := Factorial(val.input)
		assert.Equal(t, res, val.expected, val.name)
	}
}
