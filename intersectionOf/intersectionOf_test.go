package intersectionof

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectionOf(t *testing.T) {
	testCases := []struct {
		A      []int
		B      []int
		actual []int
	}{
		{
			A:      []int{1, 2, 3, 4, 5},
			B:      []int{2, 3, 5, 9},
			actual: []int{2, 3, 5},
		},
		{
			A:      []int{},
			B:      []int{1, 2, 3},
			actual: []int{},
		},
		{
			A:      []int{1, 3, 5, 7, 9},
			B:      []int{0, 2, 4, 6, 8},
			actual: []int{},
		},
		{
			A:      []int{3, 1, 2},
			B:      []int{1, 3, 2},
			actual: []int{1, 2, 3},
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, IntersectionOf(tc.A, tc.B), tc.actual)
	}
}

func TestIntersectionOf2(t *testing.T) {
	testCases := []struct {
		A      []int
		B      []int
		actual []int
	}{
		{
			A:      []int{1, 2, 3, 4, 5},
			B:      []int{2, 3, 5, 9},
			actual: []int{2, 3, 5},
		},
		{
			A:      []int{},
			B:      []int{1, 2, 3},
			actual: []int{},
		},
		{
			A:      []int{1, 3, 5, 7, 9},
			B:      []int{0, 2, 4, 6, 8},
			actual: []int{},
		},
		{
			A:      []int{3, 1, 2},
			B:      []int{1, 3, 2},
			actual: []int{1, 2, 3},
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, IntersectionOf2(tc.A, tc.B), tc.actual)
	}
}
