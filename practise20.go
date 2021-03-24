/*
 * @Description:
 * @Author: XiaoAn
 * @Date: 2021-02-19 14:37:27
 * @LastEditTime: 2021-02-19 14:43:35
 * @LastEditors: XiaoAn
 */
package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}
type outerS struct {
	b int
	c float32
	int
	innerS
}

func main() {
	outer := new(outerS)
	outer.b = 14
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 15
	outer.in2 = 19
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	fmt.Println(outer.innerS.in1)
}
