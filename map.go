package main

import "fmt"

func main() {
	var usermap = map[int]string{
		1: "小明",
		2: "小红",
		4: "",
	}
	fmt.Println(usermap[1])
	fmt.Printf("%#v\n", usermap[3])
	fmt.Printf("%#v\n", usermap[4])

	usermap[1] = "修改了"
	fmt.Println(usermap)

	delete(usermap, 1)
	fmt.Println(usermap)
}
