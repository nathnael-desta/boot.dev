package main

func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int)
	for _, name := range names {
		r := []rune(name)[0]
		_, ok := nameCounts[r]
		if !ok {
			nameCounts[r] = make(map[string]int)
		}
		nameCounts[r][name]++
	}
	return nameCounts
}
