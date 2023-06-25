package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{1, 2, 3, 4, 6, 7}

	// 返回 target 在 slice 中第一次出现的位置，
	index, b := slices.BinarySearch(s, 6)
	if b {
		// Output: 4
		fmt.Println(index)
	}
}
