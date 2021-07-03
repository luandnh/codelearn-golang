package numOfLetters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumOfLetters(t *testing.T) {
	testCases := []struct {
		str    string
		actual int
	}{
		{
			str:    "abcd",
			actual: 0,
		},
		{
			str:    "abacb",
			actual: 1,
		},
		{
			str:    "goodcodegoodlifegoodgamegg",
			actual: 5,
		},
		{
			str:    "congchanhunuithaisonnghiamenhunuoctrongnguonchayra",
			actual: 3,
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, NumOfLetters(tc.str), tc.actual)
	}
}
