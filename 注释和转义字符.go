package main

import "fmt"

// 函数外界变量一定要用var来声明,不用也不会报错
var shuxing = "这是一个变量"

// 定义常量用const
const version string = "1.0.0"

var (
	s1 string = "小花子"
	s2 string = "小mi子"
)

func hello(name string) {
	fmt.Println(name, "说你好")
}

func main() {
	//1. 先声明再赋值
	var name string
	name = "小李子"
	fmt.Println("name是", name)
	//2.声明并赋值
	var name1 string = "小刚子"
	fmt.Println("name1是", name1)

	//3.省略声明
	var name2 = "小6子"
	fmt.Println("name3是", name2)

	//3.类型推导
	name3 := "小花子"
	fmt.Println("name2是", name3)

	// 用\"来转义
	fmt.Println("\"hello world\"")

	hello(name)

	//多个变量
	var a1, a2 = 1, 2
	fmt.Println(a1, a2)

}
