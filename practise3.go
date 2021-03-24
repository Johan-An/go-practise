package main

import "fmt"
// 不使用中间变量实现变量交换
func main() {
	a := 1
	b := 2
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Println(a, b)
}
