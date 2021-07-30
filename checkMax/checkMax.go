package checkmax

func CheckMax(str string) int {
	max := 0
	for _, char := range str {
		if max < int(char) {
			max = int(char)
		}
	}
	return max
}
