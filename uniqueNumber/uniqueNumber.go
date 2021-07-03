package uniquenumber

func uniqueNumber(arr []int) int {
	tmp := make(map[int]int, len(arr))
	for _, value := range arr {
		if tmp[value] == 0 {
			tmp[value] = 1
		} else {
			tmp[value] = tmp[value] + 1
		}
	}
	for key, value := range tmp {
		if value == 1 {
			return key
		}
	}
	return 0
}
