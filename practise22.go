/*
 * @Description:
 * @Author: XiaoAn
 * @Date: 2021-02-23 17:56:22
 * @LastEditTime: 2021-02-24 15:18:19
 * @LastEditors: XiaoAn
 */
package main

import "io"

type Socket struct {
}

func (s *Socket) Write(p []byte) (n int, err error) {
	return 0, nil
}
func (s *Socket) Close() error {
	return nil
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}

func usingWriter(writer io.Writer) {
	writer.Write(nil)
}
func usingCloser(closer io.Closer) {
	closer.Close()
}
func main() {
	// 实例化Socket
	s := new(Socket)
	usingWriter(s)
	usingCloser(s)
}
