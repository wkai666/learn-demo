package main

import (
	"fmt"
	"image/gif"
	"math/rand"
	"os"
)

/*
	var 变量声明语法可以创建一个特定类型的变量，然后给变量附加一个名字，并设置变量的初始值。
	变量声明的一般语法为：var 变量名 类型 = 表达式
	其中，类型 或者 表达式 可以省略其中之一。
	如果省略的是类型信息，那么将根据初始表达式来推导变量的类型信息。
	如果初始化表达式被省略，那么将用零值初始化该变量。
	数值类型变量对应的零值是 0
	布尔类型变量对应的零值是 false
	字符串类型对应的零值是空字符串
	接口或者引用类型(包括 slice、指针、map、chan和函数) 变量对应的初始值是 nil
	数组或者结构体等聚合类型对应的零值是每个元素或者字段都是对应的该类型的零值
*/

/*
	零值初始化可以确保每个声明的变量总是有一个良好定义的值，因此在 go 语言中，不存在为初始化的变量。
	上述这个特性可以简化代码，而且可以在没有增加额外工作的前提下确保边界条件下的合理性为。

*/

/*
	在包级别声明的变量会在 main入口函数执行前完成初始化。
	局部变量将在生命语句被执行到的时候完成初始化。
*/

/*
	简短变量声明被广泛用于大部分的局部变量的声明和初始化。
	var 形式的变量声明往往用于需要显式指定变量类型的地方，或者因为变量稍后会被重新赋值而初始值无关紧要的地方。
*/

/*
	:= 是一个变量声明语句
	= 是一个变量赋值操作
*/

func main() {
	// 这段代码将打印空字符串，而不是导致错误或者不可预知的行为。
	// Go 语言的程序员应该让一些聚合类型的零值也具有意义，这样可以保证不管任何类型的变量总有一个合理有效的零值状态。
	var s string
	fmt.Println("string s is: ", s)

	// 在一个声明语句中同时声明一组变量
	var i, j, k int
	// 用一组初始化表达式声明并初始化一组变量
	// 省略每个类型的变量，则可以声明多个类型不同的变量（类型由初始化表达式推导）
	var b, f, t = true, 2.3, "four"
	fmt.Println("i is:", i)
	fmt.Println("j is:", j)
	fmt.Println("k is:", k)
	fmt.Println("b is:", b)
	fmt.Println("f is:", f)
	fmt.Println("t is:", t)
	// 初始化表达式可以是字面量或者任意表达式。

	// 一组变量也可以通过调用一个函数，由函数的多个返回值初始化
	/*
		简短变量声明中，左边的变量可能并不是刚刚全部声明。
		如果有一些已经在相同的词法域中声明过了，那么简短变量声明语句对这些已经声明过的变量就只有赋值行为了。
		简短变量声明语句中，必须至少要声明一个新的变量

	*/
	var file, err = os.Open("const.go")
	if err != nil {
		panic(err)
	}
	file.Close()

	/*
		简短变量的声明: 变量名 := 表达式
		函数内部用于声明和初始化局部变量，变量的类型根据表达式来自动推导。
	*/
	anim := gif.GIF{LoopCount: 5}
	freq := rand.Float64() * 3.0
	tt := 0.0

	fmt.Println("anim is: ", anim)
	fmt.Println("freq is: ", freq)
	fmt.Println("t is: ", tt)

	// 简短变量声明被广泛用于大部分的局部变量的声明和初始化。
	// var 形式的变量声明往往用于需要显式指定变量类型的地方，或者因为变量稍后会被重新赋值而初始值无关紧要的地方。
	ii := 100
	var boiling float64 = 100
	var names []string
	var e error
	// var p Point

	fmt.Println("ii is: ", ii)
	fmt.Println("boiling is: ", boiling)
	fmt.Println("names is: ", names)
	fmt.Println("err is: ", e)
	//fmt.Println("p is: ", p)

	/*
		和 var 形式生命语句一样，简短变量声明语句也可用来声明和初始化一组变量
		但是同时声明多个变量的方式应该限制只在可以提高代码可读性的地方使用，
		比如：for 语句的循环的初始化语句部分
	*/
	i, j = 0, 1

	/*
		多重赋值：将右边表达式的值赋给左边对应位置的各个变量
	*/
	i, j = j, i

	/*
		简短变量声明语句只有对已经在同级词法域声明过的变量才和赋值操作语句等价，
		如果变量是在外部词法域声明的，那么简短变量声明语句将会在当前词法域重新声明一个新的变量。
	*/

	/*
		指针：保存了对应类型值的内存空间。
		普通变量在生命语句创建时被绑定到一个变量名，比如叫x的变量，但是还有很多变量始终以表达式的方式引入
		例如：x[i] 或者x.f变量。所有这些表达式一般都是读取一个变量的值，除非他们是出现在赋值语句的左边，这时候是给对应变量赋予一个新的值。

		一个指针的值是另一个变量的地址。一个指针对应变量在内存中的存储位置。并不是每个值都会有一个内存地址，
		但是对于每一个变量必然有对应的内存地址。通过指针，我们可以直接读取或者更新对应变量的值，而不需要知道改变量的名字（如果变量有名字的话）。

		如果用var x int 声明语句声明一个 x 变量，那么 &x表达式（即x变量的内存地址）讲产生一个指向该整数变量的指针，
		指针对应的类型数据是 *int，指针被称之为"指向int类型的指针"。如果指针名字为p，那么可以说"p指针指向变量x"，或者说"p指针保存了x变量的内存地址"
		同时，*p 表达式对应 p指针指向变量的值。一般 *p表达式读取指针指向变量的值，这里为 int 类型的值，
		同时因为 *p对应一个变量，所以该表达式也可以出现在赋值语句的左边，表示更新指针所指向变量的值。

	*/
	x := 1
	p := &x
	fmt.Println(*p)

	*p = 2
	fmt.Println(x)

	/*
		对于聚合类型每个成员，比如结构体的每个字段，或者是数组的每个元素，也都是对应一个变量，因此可以被取地址。

		变量有时候被称为可寻址的值。即使变量由表达式临时生成，那么表达式也必须能接受 & 取地址操作。
		任何类型指针的零值都是 nil。如果 p 指向某个有效变量，那么 p != nil 测试为真。
		指针之间也是可以进行相等测试的，只有当他们指向同一个变量或全部是 nil 时才相等。
	*/
	var xx, y int
	fmt.Println(&xx == &xx, &xx == &y, &xx == nil)

	/*
		Go 语言中，返回函数中局部变量的地址也是安全的。例如下面的代码，调用f函数时，创建局部变量 v，
		在局部变量地址被返回之后依然有效，因为指针p依然引用这个变量。
	*/
	var pp = ff()

	fmt.Println(pp)
	fmt.Println(ff() == ff())
}

func ff() *int {
	v := 1
	return &v
}
