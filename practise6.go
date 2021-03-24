/**
 * @name: 解析命令输入的参数
 * @Autor: Xiaoan
 * @param {*}
 * @return {*}
 */
package main

import (
	"flag"
	"fmt"
)

// flag.Type(name, defValue, usage)，例如：go run .\practise6.go -mode common
var mode = flag.String("mode", "esay", "process model")

func main() {
	flag.Parse()
	fmt.Println(*mode)
}
