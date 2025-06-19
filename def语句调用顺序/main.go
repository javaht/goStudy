package main

import "fmt"

func main() {
	//写入defer语句，在函数结束前执行
	//defer 语句的执行顺序是后进先出（LIFO  栈类型
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	fmt.Println("main")

}
