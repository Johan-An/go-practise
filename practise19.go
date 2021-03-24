/*
 * @Description:
 * @Author: XiaoAn
 * @Date: 2021-02-04 18:54:14
 * @LastEditTime: 2021-02-04 18:57:30
 * @LastEditors: XiaoAn
 */
package main

import "fmt"

type People struct {
	name string
	sex  string
}

func main() {
	people := People{
		name: "jack",
		sex:  "man",
	}
	fmt.Println(people)
}
