package main

import "fmt"

/*
	1、数组及切片反转的实现
	2、slice 作为参数传递时，是指针传递，会改变原来切片的值
*/

func main() {
	// 切片的反转
	s := []int{3, 5, 6, 9, 12, 14}
	fmt.Printf("origin s is: %+v\n", s)
	reverse(s) // 切片的反转
	fmt.Printf("reverse s is: %+v\n", s)

	// 数组的反转
	a := [...]int{1, 2, 4, 5, 7, 9}
	fmt.Printf("a before reverse is: %+v\n", a)
	reverse(a[:])
	fmt.Printf("after reverse a is: %+v\n", a)
}

// reverse 数组、切片反转函数
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		// 交换位置
		s[i], s[j] = s[j], s[i]
	}
}
