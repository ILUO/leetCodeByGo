package main

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	isCycle := false
	for !isCycle && start < len(gas){
		remain := 0
		for i := start;;{
			remain += gas[i] - cost[i]
			if remain < 0{
				start = i + 1
				break
			}
			i = (i+1) % len(gas)
			if i == start{
				return start
			}
			if i == 0{
				isCycle = true
			}
		}
	}
	return -1
}
