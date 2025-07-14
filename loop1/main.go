package main

func maxMessages(thresh int) int {
	totalCost := 0
	for i := 1; ; i++ {
		totalCost += 100 + i - 1
		if totalCost > thresh {
			return i - 1
		} 
	}
}
