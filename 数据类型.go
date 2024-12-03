package main

import "fmt"

func main() {

	var age = 12
	fmt.Println(age)
	//定义了一个无符号8位整型变量 u8 并将其初始化为255
	var u8 uint8 = 255
	fmt.Println(u8)
	// int8
	var i8 int8 = -128
	fmt.Println(i8)

	fmt.Println("=========================")

	//浮点型 float32  float64

	//字符型 byte(单字节字符) 和rune(多字节字符)

	//字符串类型  双引号
	fmt.Println("周州\t知道")
	fmt.Println("'周州'知道")
	fmt.Println("\"周州\"知道")

	//多行字符
	fmt.Printf(`
						真
						你
						吗
						`)

	//布尔类型 默认false 不允许整形强转为布尔类型 无法参与算术运算

	//零值问题 int就是0  布尔是flase 字符串是"" 浮点型是0.0  byte是0

}
