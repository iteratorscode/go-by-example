package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// 声明变量
	var a = "initial"
	fmt.Println(a)

	var b, c = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// 另一种方式声明变量，类型推断
	f := "apple"
	fmt.Println(f)
}
