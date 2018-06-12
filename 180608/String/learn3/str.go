package main

import (
	"fmt"
	"strings"
)
//初始位置 (0, 0) 处有一个机器人。给出它的一系列动作，判断这个机器人的移动路线是否回到起点
//移动顺序由一个字符串表示。每一个动作都是由一个字符来表示的。机器人有效的动作有 R（右），L（左），U（上）和 D（下）
//输出应为 true 或 false，表示机器人移动路线是否成圈
func main() {
	fmt.Println(judgeCircle2("UD"))
	fmt.Println(judgeCircle2("LABCL"))
}

func judgeCircle(moves string) bool {
	x := 0
	y := 0
	for _, ch := range moves {
		switch ch {
		case 'L':
			x++
		case 'R':
			x--
		case 'U':
			y++
		case 'D':
			y--
		default:
			break
		}
	}
	return x == 0 && y == 0
}


//大神之作
func judgeCircle2(moves string) bool {
	if len(moves)%2 == 1{
		return false
	}

	if(strings.Count(moves, "U") == strings.Count(moves, "D") &&
		strings.Count(moves, "L") == strings.Count(moves, "R")){
		return true
	}

	return false
}