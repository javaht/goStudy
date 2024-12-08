package main

import "fmt"

func main() {

	var age int
	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)

	switch {
	case age <= 0:
		fmt.Println("还没出生")
	case age > 0 && age < 18:
		fmt.Println("未成年")
	case age <= 35:
		fmt.Println("青年")
	default:
		fmt.Println("老年")
	}
}
