package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	s := strconv.Itoa(1)
    return countAndSayCore(n,0,s)
}

func countAndSayCore(n,k int,s string) string{
	if n-1 == k{
		return s
	}
	var a []string
	for i,j := 0,0;i <= len(s);i++{
		if i == len(s) || s[j] != s[i]{
			a = append(a, s[j:i])
			j = i
		}
	}
	res := bytes.Buffer{}
	for _,str := range a{
		l := len(str)
		res.WriteString(strconv.Itoa(l)+str[0:1])
	}
	return countAndSayCore(n,k+1,res.String())
}

func main(){
	fmt.Println(countAndSay(4))
}