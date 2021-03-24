/*
 * @Description:break练习
 * @Author: XiaoAn
 * @Date: 2021-01-27 17:26:56
 * @LastEditTime: 2021-01-27 17:37:53
 * @LastEditors: XiaoAn
 */
package main

import "fmt"

func main() {
outLoop:
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			fmt.Println(x, y, "continue...")
			switch y {
			case 3:
				fmt.Println(x, y)
				// break outLoop
				continue outLoop
			case 4:
				fmt.Println(x, y)
				// break outLoop
				continue outLoop
			}
		}
	}
	fmt.Println("nice")
}
