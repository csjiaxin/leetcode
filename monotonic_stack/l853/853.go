package l853

import "sort"

// https://leetcode.com/problems/car-fleet/
// Embedded subtemplate: see README.md for its invariant and tie policy.
type fleetCar struct{ position, speed int }

func countFleets(target int, position, speed []int) int {
	cars := make([]fleetCar, len(position))
	for i, p := range position {
		cars[i] = fleetCar{p, speed[i]}
	}
	sort.Slice(cars, func(i, j int) bool { return cars[i].position > cars[j].position })
	fleets := 0
	lastArrival := -1.0
	for _, car := range cars {
		arrival := float64(target-car.position) / float64(car.speed)
		if arrival > lastArrival {
			fleets++
			lastArrival = arrival
		}
	}
	return fleets
}

func carFleet(target int, position []int, speed []int) int {
	return countFleets(target, position, speed)
}
