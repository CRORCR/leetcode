package main

import (
	"fmt"
	"strings"
)

//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

//示例 1:
//输入: "Let's take LeetCode contest"
//输出: "s'teL ekat edoCteeL tsetnoc"
//注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
func main() {
	fmt.Println(reverseWords2("Let's take LeetCode contest"))
}


//思路:先按照空格分割,然后根据每个单词去翻转,翻转的结果存入arr切片中,最后把切片使用空格拼接,最后去除收尾空格
func reverseWords(s string) string {
	split := strings.Split(s, " ")
	var arr=make([]string,len(split))
	for _,v:=range split{
		word:=[]rune(v)
		l:=len(word)/2
		for i:=0;i<l;i++ {
			word[i],word[len(word)-1-i]=word[len(word)-1-i],word[i]
		}
		arr=append(arr,string(word))
	}

	return strings.TrimSpace(strings.Join(arr, " "))
}

//大神之作
//直接把字符串转换成数组
//当遇到空格,就把从标志点spos到空格前一个字符去翻转
//然后再重置标识点spos
func reverseWords2(s string) string {
	buf := []byte(s)
	spos := 0

	for i := 0; i <= len(buf); i++ {
		if i == len(buf) || buf[i] == ' ' {
			for m, n := spos, i-1; m < n; m, n = m+1, n-1 {
				buf[m], buf[n] = buf[n], buf[m]
			}
			spos = i + 1
			continue
		}
	}

	return string(buf)
}