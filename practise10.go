/*
 * @Author: your name
 * @Date: 2021-01-22 11:53:04
 * @LastEditTime: 2021-01-27 16:17:49
 * @LastEditors: XiaoAn
 * @Description: In User Settings Edit
 * @FilePath: \go\practise10.go
 */
package main

import "fmt"

/**
 * @name: 九九乘法表
 * @Autor: Xiaoan
 * @param {*}
 * @return {*}
 */
func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
}
