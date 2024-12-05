package main

import (
	"fmt"
	"sort"
)

func main() {

	//这个是数组
	//var nameList [3]string = [3]string{"小明", "小红", "小刚"}
	//fmt.Println("nameList是", nameList)
	//fmt.Println("nameList[0]是", nameList[0])
	//fmt.Println("nameList[1]是", nameList[1])
	//fmt.Println("nameList[2]是", nameList[2])
	//fmt.Println(nameList[len(nameList)-1])
	//nameList[0] = "换个呗"
	//fmt.Print(nameList)

	//这个是切片 更习惯叫列表
	//定义一个字符串切片
	//var nameList []string   //或者 nameList := make([]string, 0)
	//nameList = append(nameList, "小明")
	//nameList = append(nameList, "小红")
	//nameList = append(nameList, "小刚")
	//fmt.Println(nameList)

	//如果定义切片没有赋值会报错
	// make([]type,length,容量)
	//nameList := make([]string, 0)
	//fmt.Println(nameList == nil)

	//array := [3]int{1, 2, 3}
	//slices := array[:]
	//slices := array[0:2] //从0到1不包括2
	//fmt.Println(slices)

	//切片排序
	var ints = []int{4, 3, 7, 9, 0}
	sort.Ints(ints)
	fmt.Println(ints)
	//倒序
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints)
}
