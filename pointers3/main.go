package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line
func getMessageText(a *Analytics, m Message) {
	switch m.Success {
	case true:
		a.MessagesSucceeded++
		a.MessagesTotal++
	case false:
		a.MessagesFailed++
		a.MessagesTotal++
	}
	
} 
// ?
