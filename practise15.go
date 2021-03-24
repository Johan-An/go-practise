/*
 * @Description:
 * @Author: XiaoAn
 * @Date: 2021-02-01 14:20:29
 * @LastEditTime: 2021-02-01 14:29:29
 * @LastEditors: XiaoAn
 */
package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "", "skill to perform")

/**
 * @name: 匿名函数实现操作封装
 * @Autor: Xiaoan
 * @param {*}
 * @return {*}
 */
func main() {
	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angle fly")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		fmt.Println(ok)
		f()
	} else {
		fmt.Println("skill not found")
	}
}
