/*
 * @Author: your name
 * @Date: 2021-01-19 11:51:53
 * @LastEditTime: 2021-01-19 18:22:39
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \go-work\go\practise7.go
 */
package main

import "fmt"

func main() {
	var a = make([]int, 2, 3)
	fmt.Println(cap(a))
	a = append(a, 4)
	fmt.Println(a)
	a = append([]int{1}, a...)
	fmt.Println(a)
	a = append([]int{-1, 1}, a...)
	fmt.Println(a)
}
