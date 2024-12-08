package main

import "fmt"

func main() {
	//var sum int
	//for i := 1; i <= 100; i++ {
	//	sum = sum + i
	//}
	//println(sum)

	//for循环的五种形式()

	//1.死循环
	//for i := 1; true; i++ {
	//	fmt.Println(time.Now())
	//	time.Sleep(2 * time.Second)
	//}

	//2.省略
	//for true {
	//	fmt.Println(time.Now())
	//	time.Sleep(2 * time.Second)
	//}

	//3.while循环
	//var sum int
	//var i int = 1
	//for i <= 100 {
	//	sum = sum + i
	//	i++
	//}
	//fmt.Println(sum)

	//4.do while模式
	//var sum int
	//var i int = 1
	//for {
	//	sum = sum + i
	//	i++
	//	if i == 100 {
	//		break
	//	}
	//}

	//5.利用切片方式
	var list = []string{"周周", "知道"}
	for index, value := range list {
		fmt.Println(index, value)
	}

}
