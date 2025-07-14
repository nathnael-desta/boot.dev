package main

func countConnections(groupSize int) int {
	totalCount := 0
	for i := 1; i < groupSize + 1; i++ {
		totalCount += i - 1
	}
	return totalCount
}
