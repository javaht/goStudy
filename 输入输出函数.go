package main

import "fmt"

func main() {
	//  格式化输出
	//  %s代表字符串
	//  %d代表整数
	//  %f代表浮点数
	//  %T用来打印类型
	//  %v代表任意的
	fmt.Printf("%s  哇,你好美", "峰峰") //输出    峰峰  哇,你好美

	//以下是输入
	fmt.Println("请输入你的名字:")
	var name string
	fmt.Scanln(&name)
	fmt.Println("你的名字是:", name)
}
