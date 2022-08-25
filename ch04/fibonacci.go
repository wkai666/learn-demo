package main

import "fmt"

func main() {
	// 闭包函数实现斐波那契数列求和
	f := fibonacciC() // 返回一个闭包函数
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// 递归实现斐波那契数列求和
func fibonacci(n int) int {
	if n <= 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

// 闭包实现斐波那契数列求和
func fibonacciC() func() int {
	back1, back2 := 0, 1
	return func() int {
		back1, back2 = back2, (back1 + back2)
		return back1
	}
}
