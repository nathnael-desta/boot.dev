package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	var messageCosts [3]int
	totalCost := 0 
	for message := range messages {
		totalCost += len(messages[message])
		messageCosts[message] = totalCost
	}
	return messages, messageCosts
}
