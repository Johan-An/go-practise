package main

import "net/http"
// go搭建web服务
func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8087", nil)
}
