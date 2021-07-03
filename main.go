package main

import (
	intersectionof "codelearn-golang/intersectionOf"
	"fmt"
)

func main() {
	A := []int{1, 2, 3, 4, 5}
	B := []int{2, 3, 5, 9}
	fmt.Print(intersectionof.IntersectionOf2(A, B))
}
