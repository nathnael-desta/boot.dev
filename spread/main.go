package main

func sum(nums ...int) int {
	totalCost := 0
	for i := 0 ; i < len(nums); i++ {
		totalCost += nums[i]
	}
	return totalCost
}
