/*
 * @Description:
 * @Author: XiaoAn
 * @Date: 2021-03-08 15:26:10
 * @LastEditTime: 2021-03-08 15:38:22
 * @LastEditors: XiaoAn
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index) // index 为向url发送请求时，调用的函数
	log.Fatal(http.ListenAndServe("localhost:8088", nil))
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "c语言中文网")
}
