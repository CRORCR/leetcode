package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now().Minute()

	_ := numJewelsInStones("aAasfcasf", "aAAbbb")
	end := time.Now().Minute()
	fmt.Println(end-start)
}

//
//我考虑的是strings包下有一个count计数的函数,可以拿到J的每个字符去统计
func numJewelsInStones(J string, S string) int {
	var count int
	for i:=0;i<len(J);i++{
		u := string(J[i])
		count += strings.Count(S, u)
	}
	return count
}

//大神代码,使用 2个range,第一个字符==第二个字符就++
//快在使用了range
func numJewelsInStones2(J string, S string) int {
	a:=0
	for _,v:=range J{
		for _,v1:=range S{
			if v==v1{
				a++
			}
		}
	}
	return a
}