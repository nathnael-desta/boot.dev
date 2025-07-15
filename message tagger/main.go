package main

import (
	"strings"

)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i, message := range messages {
		messagetTags := tagger(message)
		if messagetTags != nil {
			messages[i].tags = messagetTags
		}
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	content := strings.ToLower(msg.content)
	if (strings.Contains(content, "urgent")) {
		tags = append(tags, "Urgent")
	}
	if (strings.Contains(content, "sale")) {
		tags = append(tags, "Promo")
	}
	return tags
}
