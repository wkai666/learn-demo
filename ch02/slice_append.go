package main

import "fmt"

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap = %d\t%v\n", i, cap(y), y)
		x = y
	}
}

/*
	每次调用 appendInt 函数，必须先检测 slice 底层数组是否有足够的容量来保存新添加的元素。
	如果有足够的空间的话，直接扩展slice(依然在原有的底层数组之上)，将新添加的 y 元素复制到新
	扩展的空间上，并返回 slice。因此，输入的 x 和输出的 z 共享相同的底层数组。

	如果没有足够的增长空间的话，appendInt 函数则会先分配一个足够大的 slice 用于保存新的结果，
	先将输入的 x 复制到新的空间，然后添加 y 元素。结果 z 和输入的 x 引用的将是不同的底层数组。

	内置的 copy 函数可以很方便的将一个 slice复制到另一个想同类型的 slice。

	为了提高内存使用效率，新分配的数组一般略大于保存 x 和 y 所需要的最低大小。
	通过在每次扩展数组时，直接将长度翻倍从而避免了多次内存分配，也确保了添加单个元素操作的平均时间是一个常数。
*/

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1 // 新切片的长度
	if zlen <= cap(x) {
		// 原切片容量足够
		z = x[:zlen]
	} else {
		// 原切片容量不足时，对切片进行 2 倍扩容
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	// 追加元素到切片中
	z[len(x)] = y
	return z
}
