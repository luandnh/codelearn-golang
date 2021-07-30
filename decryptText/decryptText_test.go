package decrypttext

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDecryptText(t *testing.T) {
	testCases := []struct {
		str    string
		actual string
	}{
		{
			str:    "cqtximmdq",
			actual: "codelearn",
		},
		{
			str:    "ibttlprimfymqlpgeftwu",
			actual: "itsasecrettoeverybody",
		},
		{
			str:    "ftnexvoky",
			actual: "fourtytwo",
		},
		{
			str:    "taevzhzmashvjw",
			actual: "thereisnospoon",
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, DecryptText(tc.str), tc.actual)
	}
}
