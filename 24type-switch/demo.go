/*
 * @Description:
 * @Author: XiaoAn
 * @Date: 2021-03-08 11:29:04
 * @LastEditTime: 2021-03-08 12:08:17
 * @LastEditors: XiaoAn
 */
package main

import "fmt"

// 电子支付
type Alipay struct {
}

// 为Alipay 添加CanUseFaceID方法，表示支付方式支持刷量支付
func (a *Alipay) CanUseFaceID() {

}

// 现金支付
type Cash struct {
}

// 为现金方式添加Stolen方法
func (a *Cash) stolen() {
}

// 具备刷脸支付的接口
type CantainCanUseFaceID interface {
	CanUseFaceID()
}

// 具备可被偷窃的性值
type ContainStolen interface {
	stolen()
}

func print(payMethod interface{}) {
	switch payMethod.(type) {
	case CantainCanUseFaceID:
		fmt.Println("%T can use faceid\n", payMethod)
	case ContainStolen:
		fmt.Println("%T may be stolen\n", payMethod)
	}

	// fmt.Println(payMethod.(type))
}
func main() {
	print(new(Alipay))
	print(new(Cash))
}
