package main

import "fmt"

func divide(dividend int, divisor int) int {
	IntMin := -(1<<31)
	if divisor == -1 && dividend == IntMin{
		return 1<<31 - 1
	}
	isN := false
	var dd int64
	var dr int64
	if dividend < 0{
		isN = !isN
		dd = int64(-dividend)
	}else{
		dd = int64(dividend)
	}
	if divisor < 0{
		isN = !isN
		dr = int64(-divisor)
	}else{
		dr = int64(divisor)
	}
	res := divideCore(dd,dr)
	if isN{
		return -res
	}
	return res
}

func divideCore(dd ,dr int64) int{
	if dd < dr{
		return 0
	}
	if dd == dr{
		return 1
	}
	count := 1
	tempDr := dr
	for ;dd >= dr + dr;{
		count = count + count
		dr = dr + dr
	}
	return count + divideCore(dd - dr,tempDr)
}

func main(){
	fmt.Println(divide(-2147483648,-1))
}