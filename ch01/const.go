package main

import "fmt"

// go 语言的 4 种声明类型：const、var、func、type
// const 	常量
// var 		变量
// func     函数的声明
// type		类型的声明

// 包一级声明的常量或者变量，可在整个包对应的整个源文件中访问，而不仅仅是在其声明语句所在的源文件中访问。
const boilingF = 212.0

// 函数的声明，是由函数名、参数列表、可选的返回值列表和包含函数定义的函数体组成。
// 如果没有返回值，那么返回列表是省略的。
// 执行函数从函数的第一个语句开始，依次顺序执行直到遇到 return 语句，如果没有返回语句，则执行到函数末尾，然后返回到函数调用者。
func main() {
	// 局部声明的变量仅可在函数内部很小的范围内被访问
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g ℉ or %g ℃\n", f, c)
}
