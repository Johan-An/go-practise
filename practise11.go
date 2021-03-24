/*
 * @Description:
 * @Author: XiaoAn
 * @Date: 2021-01-27 16:48:46
 * @LastEditTime: 2021-01-27 17:03:46
 * @LastEditors: XiaoAn
 */
package main

import "fmt"

/**switch 结构练习
 * @name:
 * @Autor: Xiaoan
 * @param {*}
 * @return {*}
 */
func main() {
	var a string
	a = "Mon"
	switch a {
	case "Mon":
		fmt.Println("周一")
		fallthrough
	case "Tue":
		fmt.Println("周二")
	case "Wed":
		fmt.Println("周三")
	}
}
