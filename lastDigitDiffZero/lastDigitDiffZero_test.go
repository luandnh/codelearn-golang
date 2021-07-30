package lastdigitdiffzero

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLastDigitDiffZero(t *testing.T) {
	testCases := []struct {
		str    int
		actual int
	}{
		{
			str:    5,
			actual: 2,
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, lastDigitDiffZero(tc.str), tc.actual)
	}
}
