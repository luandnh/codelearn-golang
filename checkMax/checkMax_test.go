package checkmax

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCheckMax(t *testing.T) {
	testCases := []struct {
		str    string
		actual int
	}{
		{
			str:    "abc",
			actual: 99,
		},
		{
			str:    "fdhfghfgh",
			actual: 104,
		},
		{
			str:    "slfjlJKHKJkk",
			actual: 115,
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, CheckMax(tc.str), tc.actual)
	}
}
