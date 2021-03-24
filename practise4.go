package main

import "fmt"

var c = 3
var d = 6

// 局部变量和全局变量练习
func change() {
	c = 4
	d := 8
	fmt.Println(c, d)
}
func main() {
	change()
	fmt.Println(c, d)
}
