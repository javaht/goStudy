package main

import "fmt"

func main() {
	var arr [5]int
	arr10 := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr10); i++ {
		fmt.Println(arr[i])
	}
	//遍历数组 索引和值
	for index, value := range arr10 {
		fmt.Println(index, value)
	}
}
