package main

import "fmt"

func romanToInt(s string) int {
	mToV := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0
	for i := 0;i < len(s);i++ {
		if i != len(s)-1{
			if s[i:i+2] == "IV"{
				res += 4
				i++
				continue
			}
			if s[i:i+2] == "IX"{
				res += 9
				i++
				continue
			}
			if s[i:i+2] == "XL"{
				res += 40
				i++
				continue
			}
			if s[i:i+2] == "XC"{
				res += 90
				i++
				continue
			}
			if s[i:i+2] == "CD"{
				res += 400
				i++
				continue
			}
			if s[i:i+2] == "CM"{
				res += 900
				i++
				continue
			}
		}
		res += mToV[s[i:i+1]]
	}
	return res
}

func main(){
	fmt.Println(romanToInt("MCMXCIV"))
}