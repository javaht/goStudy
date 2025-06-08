package main

import "fmt"

//Readme.md
//  局部变量的声明
//	全局变量的声明
//	多变量的声明

// 声明全局变量1、2、3
var ga int = 1
var gb = 200

//如果是方法四的话来声明全局变量的话 会报错
//短变量声明语法只能在函数内部使用，不能用于全局变量的声明。
// gc := 300

func main() {
	//1. var name string = "xiaozhou"
	// fmt.Println("name是：", name)

	//2/当不写默认值的时候 默认值是0
	// var name int
	// fmt.Println("name是：", name)

	//3.当不写类型的时候 会自动推倒出类型
	// var c = 100
	// fmt.Println("c是：", c)
	//打印出类型
	// fmt.Printf("c的类型是：%T\n", c)

	//4/最常用的方法 省去var关键字 直接自动匹配      :=表示既初始化又赋值
	//在生命全局变量的时候 第四种和前几种都不一样
	// e := 100
	// fmt.Println("e是：", e)
	// fmt.Printf("e的类型是：%T\n", e)

	// fmt.Println(ga, gb)
	// fmt.Println(gc)

	//5. 声明多个变量
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// var (
	// 	name   string  = "xiaozhou"
	// 	age    int     = 25
	// 	height float64 = 175.5
	// )
	// fmt.Println(name, age, height)

	var name, age, isStudent = "xiaozhou", 25, true
	fmt.Println(name, age, isStudent)
}
