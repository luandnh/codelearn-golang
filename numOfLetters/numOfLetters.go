package numOfLetters

import (
	"strings"
)

func NumOfLetters(str string) int {
	return strings.Count(str, string(str[0])) - 1
}
