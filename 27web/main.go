/*
 * @Description:
 * @Author: XiaoAn
 * @Date: 2021-03-08 15:43:11
 * @LastEditTime: 2021-03-08 15:43:13
 * @LastEditors: XiaoAn
 */
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 在后面加上index, 来指定访问路径
	http.HandleFunc("/index", index)
	log.Fatal(http.ListenAndServe("localhost:8089", nil))
}

func index(w http.ResponseWriter, r *http.Request){
	content, _:= ioutil.ReadFile("./index.html")
	w.Write(content)
}