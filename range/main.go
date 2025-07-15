package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i := 0; i < len(msg); i++ {
		for b := 0; b < len(badWords); b++ {
			if msg[i] == badWords[b] {
				return i
			}
		}
	} 
	return -1
}
