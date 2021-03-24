/*
 * @Description:
 * @Author: XiaoAn
 * @Date: 2021-02-04 17:56:46
 * @LastEditTime: 2021-02-04 18:33:02
 * @LastEditors: XiaoAn
 */
package main

import "fmt"

type Point struct {
	x int
	y int
}
type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}
type Command struct {
	Name    string
	Var     *int
	Comment string
}

var p Point

func main() {
	p.x = 10
	p.y = 20
	fmt.Println(p)
	tank := new(Player)
	tank.Name = "cannon"
	tank.HealthPoint = 300
	fmt.Println(tank)
	var version int = 1
	cmd := newCommand("version", &version, "show version")
	fmt.Println(cmd)
}

func newCommand(name string, varref *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     varref,
		Comment: comment,
	}
}
