/*
 * @Description:
 * @Author: XiaoAn
 * @Date: 2021-03-08 11:29:04
 * @LastEditTime: 2021-03-08 15:04:32
 * @LastEditors: XiaoAn
 */
package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, dualError{Num: f}
	}
	return math.Sqrt(f), nil
}

// 自定义错误
type dualError struct {
	Num     float64
	Problem string
}

func (e dualError) Error() string {
	return fmt.Sprintf("Error!!!,because \"%f\" is a negative number", e.Num)
}
func main() {
	result, err := Sqrt(-4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
