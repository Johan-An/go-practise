/*
 * @Description:
 * @Author: XiaoAn
 * @Date: 2021-01-27 18:18:45
 * @LastEditTime: 2021-01-28 11:52:53
 * @LastEditors: XiaoAn
 */
package main

import (
	"fmt"
	"math"
)

/**
 * @name: 函数声明练习
 * @Autor: Xiaoan
 * @param {*}
 * @return {*}
 */
func main() {
	fmt.Println(add(1, 3))
	fmt.Println(minus(4, 1))
	fmt.Println(division(4.78))
}

func add(x, y int) int {
	return x + y
}

func minus(x, y int) (z int) {
	z = x - y
	return
}

func division(x float64) (interNum float64, floatNum float64) {
	interNum = math.Floor(x)
	floatNum = x - interNum
	return
}
