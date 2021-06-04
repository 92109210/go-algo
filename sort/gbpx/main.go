package main

import (
	"fmt"
)

func main() {
	a := []int{5, 4, 3, 2, 1}
	b := gbpx(a)
	fmt.Println(b)

}

// 归并排序
func gbpx(array []int) []int {
	if len(array) == 1 {
		return array
	}
	left := gbpx(array[0 : len(array)/2])
	right := gbpx(array[len(array)/2 : len(array)])
	rs := make([]int, 0, 0)
	for i, j := 0, 0; i < len(left) || j < len(right); {
		if i >= len(left) {
			rs = append(rs, right[j])
			j++
		} else if j >= len(right) {
			rs = append(rs, left[i])
			i++
		} else if left[i] <= right[j] {
			rs = append(rs, left[i])
			i++
		} else {
			rs = append(rs, right[j])
			j++
		}
	}
	return rs
}
