/*
 * @Author: your name
 * @Date: 2021-01-21 15:52:38
 * @LastEditTime: 2021-01-21 16:17:59
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \go\practise9.go
 */
package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	london, res := scene.Load("london")
	fmt.Println(london)
	fmt.Println(res)
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
