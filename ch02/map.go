package main

import (
	"fmt"
	"sort"
)

func main() {
	// 使用 make 进行初始化 map, 等价于 stu := map[string]int{}
	students := make(map[string]int)
	students["Alice"] = 21
	students["Bob"] = 23

	// 字面值的语法创建 map
	// students := map[string]int{
	// 	"Alice": 21,
	// 	"Bob":   23,
	// }

	// map 中的元素通过下标进行访问
	fmt.Println(students["Bob"])

	// 内置的 delete 函数可以删除 map 中的元素
	delete(students, "Bob")

	// Jack 不存在时将返回 0，元素在 map 中不存在将返回对应类型的零值
	students["Jack"] = students["Jack"] + 1

	// 等价于
	students["Jack"] += 1

	// 等价于
	students["Jack"]++

	// map 中元素的遍历
	for k, v := range students {
		fmt.Printf("%s\t%d\n", k, v)
	}

	// map 中的元素是无序的，需要排序时，对 key 进行显式的排序
	var names []string // 可用 names := make([]string, 0, len(students)) 代替
	for name := range students {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, students[name])
	}

	// map 的零值
	var teachers map[string]string
	fmt.Println("teachers is nil, true or false", teachers == nil)
	fmt.Println("len of teachers is 0, true or false", len(teachers) == 0)

	// map 上的大部分操作，包括查找，删除，len 和 range 都可以安全的工作在 nil 值的map 上
	// 他们的行为和一个空的 map 类似，但是向一个 nil 值的 map 读取或者存入元素则会导致 panic 异常
	teachers["english"] = "WangXiaohua"

	// 从 map 中读取元素：通过 key 作为索引下标来访问 map 中的 value
	// 如果 key 在 map 中是存在的，那么将得到与 key 对应的 value
	// 如果 key 不存在，那么将得到 value 对应类型的零值。

	// 判断元素是否存在于 map 中，可以使用 ,ok 模式，ok 为布尔值，用于报告元素是否存在于 map 中
	if _, ok := teachers["math"]; !ok {
		// 元素在 map 中不存在的处理
	}

	// map 之间比较的实现
	fmt.Printf("student is equal student, true or false %t\n", equal(students, students))
}

func equal(x, y map[string]int) bool {
	// map 元素个数的比较
	if len(x) != len(y) {
		return false
	}

	// map 中 key 和 value 的比较
	for k, xv := range x {
		if yv, ok := x[k]; !ok || xv != yv {
			return false
		}
	}

	return true
}

/*
	有时候我们需要一个map或set的key是slice类型，但是map的key必须是可比较的类型，但是slice并不满足这个条件。
	不过，我们可以通过两个步骤绕过这个限制。
	第一步，定义一个辅助函数k，将slice转为map对应的string类型的key，确保只有x和y相等时k(x) == k(y)才成立。
	然后创建一个key为string类型的map，在每次对map操作时先用k辅助函数将slice转化为string类型。
*/

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}
