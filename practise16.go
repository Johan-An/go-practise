/*
 * @Description:
 * @Author: XiaoAn
 * @Date: 2021-02-01 17:01:48
 * @LastEditTime: 2021-02-01 17:05:12
 * @LastEditors: XiaoAn
 */
package main

import "fmt"

func main() {
	myFunc(1, 2, 34, 4)
}

/**
 * @name: 可变参数练习
 * @Autor: Xiaoan
 * @param {...int} args
 * @return {*}
 */
func myFunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
