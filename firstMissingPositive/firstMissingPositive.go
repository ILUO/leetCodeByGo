package main

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] <= len(nums) {
			if nums[i] == nums[nums[i]-1] {
				break
			}
			tmp := nums[i]
			nums[i] = nums[tmp-1]
			nums[tmp-1] = tmp
		}
	}

	for i, n := range nums {
		if i != n-1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	a := []int{1, 1}
	firstMissingPositive(a)
}
