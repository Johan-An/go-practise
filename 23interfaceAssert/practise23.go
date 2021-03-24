/*
 * @Description:
 * @Author: XiaoAn
 * @Date: 2021-03-05 10:26:57
 * @LastEditTime: 2021-03-05 11:12:51
 * @LastEditors: XiaoAn
 */
package main

import "fmt"

func main() {
	var a int
	a = 10
	getType(a)
}
func getType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("type of a is int")
	case string:
		fmt.Println("type of a is string")
	case float64:
		fmt.Println("type of a is float")
	}
}
