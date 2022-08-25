package main

import "fmt"

/*
	Go 语言中 defer 和 return 解析：
	https://blog.csdn.net/qq_15371293/article/details/124187385
*/

func main() {
	fmt.Println("return: ", demo1()) // 0
	fmt.Println("return: ", demo2()) // 2

	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
	fmt.Println(f5()) // 5
	fmt.Println(f6()) // 6
}

/*
	实际上return 执行了两步操作。
	因为返回值没有命名，所以return 之前
	首先默认创建了一个临时零值变量(假设为s)作为返回值
	然后将i赋值给s,此时s的值是0。后续的操作是针对i进行的，
	所以不会影响s, 此后因为s不会更新，所以return s 不会改变
		相当于：
			var i int
			s := i
			return s
*/
func demo1() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer 2: ", i) // 2
	}()

	defer func() {
		i++
		fmt.Println("defer 1: ", i) // 1
	}()

	return i
}

/*
	因为返回值已经提前定义了，不会产生临时零值变量，
	返回值就是提前定义的变量，后续所有的操作也都是基于已经定义的变量，
	任何对于返回值变量的修改都会影响到返回值本身。

	就相当于s就是命名的变量i, 后续所有的操作都是基于
	命名变量i(s),返回值也是i, 所以每一次defer操作，
	都会更新返回值i。
*/
func demo2() (i int) {
	defer func() {
		i++
		fmt.Println("defer 2: ", i) // 2
	}()

	defer func() {
		i++
		fmt.Println("defer 1: ", i) // 1
	}()

	return
}

func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是 x，不是返回值
	}()

	return x // 5  1. 给返回值赋值  2. defer 3.真正的 return
}

func f2() (x int) {
	defer func() {
		x++
	}()

	return 5 // 6  返回值 = x, 1. 给返回值赋值，2. 执行defer  3. 真正返回
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ // 修改的是 x
	}()

	return x // 5 1. 返回值 = y = x = 5  2. defer 修改的是 x, 3.真正的返回
}

func f4() (x int) {
	defer func(x int) {
		x++ // 修改的是函数中 x 的副本
	}(x)

	return 5 // 5  返回值 = x = 5
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)

	return 5 // 5 1. x = 5, 2. defer x= 6  3.真正的返回
}

func f6() (x int) {
	defer func(x *int) *int {
		(*x)++
		return x
	}(&x)

	return 5 // 6 1. x = 5  2. defer x =6 3. 真正的返回
}
