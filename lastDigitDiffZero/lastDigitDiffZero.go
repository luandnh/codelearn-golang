package lastdigitdiffzero

import (
	"strconv"
	"strings"
)

func lastDigitDiffZero(n int) int {
	sum := 1
	for i := 2; i <= n; i++ {
		sum *= i
	}
	tmp := strconv.Itoa(sum)
	arr := strings.Split(tmp, "")
	for i := len(arr) - 1; i >= 0; i-- {
		s, _ := strconv.Atoi(arr[i])
		if s > 0 {
			return s
		}
	}
	return 0
}
