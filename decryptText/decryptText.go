package decrypttext

import (
	"strings"
)

func DecryptText(s string) string {
	var result []string
	tmp := 0
	for _, char := range s {
		i := int(char) - 97
		x := (i - tmp%26 + 26) % 26
		character := rune(x + 97)
		tmp += x
		result = append(result, string(character))
	}
	return strings.Join(result, "")
}
