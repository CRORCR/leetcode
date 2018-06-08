package main

import "fmt"

func main() {
	s := rever("hello中国")
	fmt.Println(s)
}

func rever(str string)string{
	runes := []rune(str)
	l := len(runes) / 2
	for i:=0;i<l;i++{
		runes[i],runes[len(runes)-i-1]=runes[len(runes)-i-1],runes[i]
	}
	return string(runes)
}