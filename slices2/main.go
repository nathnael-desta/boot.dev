package main

func getMessageCosts(messages []string) []float64 {
	mySlice := make([]float64, len(messages))
	for i := range messages {
		mySlice[i] = float64(len(messages[i])) * 0.01
	}
	return mySlice
}
