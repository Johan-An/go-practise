/*
 * @Description:
 * @Author: XiaoAn
 * @Date: 2021-02-19 15:22:00
 * @LastEditTime: 2021-02-19 15:25:54
 * @LastEditors: XiaoAn
 */
package main

import "fmt"

type A struct {
	a int
}
type B struct {
	a int
}
type C struct {
	A
	B
}

func main() {
	c := &C{}
	c.A.a = 15
	// c.a = 15
	fmt.Println(c.A.a)
	// fmt.Println(c.a)
}
