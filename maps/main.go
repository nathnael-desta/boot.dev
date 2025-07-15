package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return userMap, errors.New("invalid sizes")
	}
	for i, v := range names {
		userMap[v] = user{name: names[i], phoneNumber: phoneNumbers[i]}
	}
	return  userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
