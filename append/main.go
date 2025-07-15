package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	dayCost := []float64{}
	for i := 0; i < len(costs) ; i++ {
		if costs[i].day == day {
			dayCost = append(dayCost, costs[i].value)
		}
	}
	return dayCost
}
