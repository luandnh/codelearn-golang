package intersectionof

import "sort"

func IntersectionOf(A []int, B []int) []int {
	set := make([]int, 0)
	hash := make(map[interface{}]bool)
	for _, v := range A {
		hash[v] = true
	}
	for _, v := range B {
		if _, found := hash[v]; found {
			set = append(set, v)
		}
		hash[v] = true
	}
	sort.Slice(set, func(i, j int) bool {
		return set[i] < set[j]
	})
	return set
}
func IntersectionOf2(A []int, B []int) []int {
	set := make([]int, 0)
	for i := 0; i < len(A); i++ {
		if contains(B, A[i]) {
			set = append(set, A[i])
		}
	}
	sort.Slice(set, func(i, j int) bool {
		return set[i] < set[j]
	})
	return set
}
func contains(a []int, e int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == e {
			return true
		}
	}
	return false
}
