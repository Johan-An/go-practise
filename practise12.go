/*
 * @Description:
 * @Author: XiaoAn
 * @Date: 2021-01-27 17:03:57
 * @LastEditTime: 2021-01-27 17:09:17
 * @LastEditors: XiaoAn
 */
package main

import "fmt"

// goto跳转练习
func main() {
	for x := 0; x <= 10; x++ {
		for y := 0; y <= 10; y++ {
			if y == 2 {
				goto breakHere
			}
		}
	}
	return
breakHere:
	fmt.Println("done")
}
