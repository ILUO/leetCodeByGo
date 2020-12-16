package main

import "fmt"

func removeDuplicates(nums []int) int {
	count := 0
	for i,_ := range nums{
		if i != 0 && nums[i] != nums[i-1] {
			count++
		}
	}
	if len(nums) != 0{
		count++
	}
	idx := 1
	for i,_ := range nums{
		if i != 0{
			if nums[i] != nums[i-1]{
				nums[idx] = nums[i]
				idx++
			}
		}
	}
	return count
}

func main(){
	nums := []int{1,2,3,3,5,5,6,6,7}
	removeDuplicates(nums)
	fmt.Println(nums)
}
