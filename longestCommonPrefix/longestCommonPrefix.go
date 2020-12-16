package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0{
		return ""
	}
	builder := strings.Builder{}
	count := 0
	isEnd := false
	for{
		c  := ""
		for i,str := range strs{
			if count >= len(str){
				isEnd = true
				break
			}
			if i == 0{
				c = string(str[count])
			}else{
				if c != string(str[count]){
					isEnd = true
					break
				}
			}
		}
		if isEnd{
			break
		}
		builder.WriteString(c)
		count++
	}
	return builder.String()
}

func main(){
	strs := []string{"dog","racecar","car"}
	fmt.Println(longestCommonPrefix(strs))
}
