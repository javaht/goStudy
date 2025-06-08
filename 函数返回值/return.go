package main

import "fmt"

// {前方的int就是函数的返回值类型,此时是返回一个值
func test(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}

// 如果是返回多个值，匿名的
func test1(a string, b int) (string, int) {
	a = a + "   hello world"
	b = b + 100
	return a, b
}

func test2(a string, b int) (r1 int, r2 int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	//给有名称的返回值赋值
	r1 = 1000
	r2 = 2000
	return r1, r2
}

// 两个返回值都是int
func test3(a string, b int) (r1, r2 int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	//给有名称的返回值赋值
	r1 = 1000
	r2 = 2000
	return r1, r2
}

func main() {
	// c := test("xiaozhou", 200)
	// fmt.Println("c=", c)

	// a, b := test1("xiaozhou", 200)
	// fmt.Println("a=", a)
	// fmt.Println("b=", b)

	// r1, r2 := test2("xiaozhou", 200)
	// fmt.Println("r1=", r1)
	// fmt.Println("r2=", r2)

}
