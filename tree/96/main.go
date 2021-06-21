package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Nanosecond()
	fmt.Println(numTrees(19))
	t1 := time.Now().Nanosecond()
	fmt.Println(numTrees2(19))
	t2 := time.Now().Nanosecond()
	fmt.Println(t1 - t)
	fmt.Println(t2 - t1)
}

// 可以通过，但是通过递归的方法实现效率不高，
// numTrees(i)这一步骤会产生大量的重复计算，类似于斐波那契数列
// 优化：可以定义一个全局map对计算过的数据进行保存
func numTrees(n int) int {
	if n == 0 {
		return 1
	}
	rs := 0
	for i := 0; i < n; i++ {
		left := numTrees(i)
		right := numTrees(n - i - 1)
		rs = rs + left*right
	}
	return rs
}

// 迭代
func numTrees2(n int) int {
	// 数组大小为n+1，保存0-n的数量
	slice := make([]int, n+1, n+1)
	// 从小端开始计算，类似于动态规划
	slice[0] = 1
	for i := 1; i < len(slice); i++ {
		rs := 0
		for j := 0; j < i; j++ {
			rs = rs + slice[j]*slice[i-j-1]
		}
		slice[i] = rs
	}
	return slice[n]
}
