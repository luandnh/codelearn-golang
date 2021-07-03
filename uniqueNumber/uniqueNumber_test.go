package uniquenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueNumber(t *testing.T) {
	testCases := []struct {
		exp    []int
		actual int
	}{
		{
			exp:    []int{19, 17, 19, 68, 68},
			actual: 17,
		},
		{
			exp:    []int{34, 76, 76},
			actual: 34,
		},
		{
			exp:    []int{46, 46, 11, 11, 59, 59, 55, 55, 35},
			actual: 35,
		},
		{
			exp:    []int{19, 1, 41, 41, 94, 1, 94, 80, 19},
			actual: 80,
		},
		{
			exp:    []int{86, 51, 51, 72, 57, 57, 78, 78, 86},
			actual: 72,
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, uniqueNumber(tc.exp), tc.actual)
	}
}
