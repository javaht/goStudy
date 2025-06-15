package main

import "fmt"

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa // temp = main::a
	*pa = *pb  // main::a = main::b
	*pb = temp // main::b = temp
}

func main() {
	var a int = 10
	var b int = 20

	swap(&a, &b)

	fmt.Println("a = ", a, " b = ", b)
}

//解释过程
// &a 表示取变量 a 的地址，即指针变量 pa 指向变量 a 的地址。
// &b 表示取变量 b 的地址，即指针变量 pb 指向变量 b 的地址
// *pa 表示取指针 pa 所指向的变量 a 的值，即 a 的值。
// *pb 表示取指针 pb 所指向的变量 b 的值，即 b 的值。

// temp = *pa 表示将变量 temp 的值改为指针 pa 所指向的变量 a 的值。也就是说，将变量 temp 的值改为变量 a 的值。
// *pa = *pb 表示将指针 pa 所指向的变量 a 的值改为指针 pb 所指向的变量 b 的值。也就是说，将变量 a 的值改为变量 b 的值。
// *pb = temp 表示将指针 pb 所指向的变量 b 的值改为变量 temp 的值。也就是说，将变量 b 的值改为变量 temp 的值。
