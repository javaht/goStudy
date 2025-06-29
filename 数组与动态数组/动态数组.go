package main

import "fmt"

func printArray(myarray []int) {
	for _, value := range myarray {
		//下划线表示忽略某个变量
		fmt.Println(value)
	}
	myarray[0] = 100 //修改数组的值
	fmt.Println(myarray)
}

func main() {

	//[] 代表动态数组
	myarray := []int{1, 2, 3, 4, 5} //动态数据  切片slice
	//for index, value := range myarray {
	//	fmt.Println(index, value)
	//}

	printArray(myarray)
}
