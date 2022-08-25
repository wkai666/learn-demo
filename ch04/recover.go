package main

import "fmt"

/*
	输出结果：
	c  a  d  something is err  e
*/
func main() {
	fmt.Println("c")

	defer func() {
		fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("e")
	}()

	f()
	fmt.Println("f")
}

func f() {
	fmt.Println("a")
	panic("something is err")
	fmt.Println("b") // 这里开始，后续代码不会被执行
	fmt.Println("f")
}
