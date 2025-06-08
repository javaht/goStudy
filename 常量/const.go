package main

import "fmt"

// 用定义常量的关键字const来定义枚举
// const (
//
//	//可以在const()添加一个关键字iota，每行的iota都会累加1，第一行的iota默认值是0
//	beijing = iota
//	shanghai
//	shenzhen
//
// )
const (
	a, b = iota + 1, iota + 2
	// c d的公示和第一行的公式是相同的
	c, d // c = iota + 1, d = iota + 2

	e, f = iota * 2, iota * 3
	g, h
)

func main() {
	// 常量定义后无法修改
	// 	const aaa = 1
	// 	fmt.Println(aaa)
	// 	fmt.Println(shanghai)
	//

	fmt.Println(a, b)
	fmt.Println(c, d)

}
